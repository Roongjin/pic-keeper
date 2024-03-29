package user

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"

	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/Roongkun/software-eng-ii/internal/third-party/s3utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func hashEmail(email string) string {
	hasher := sha256.New()
	hasher.Write([]byte(email))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func validateImage(file io.ReadSeeker) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read file header: %w", err)
	}
	// Seek to the start of the file after reading the header
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to reset file read pointer: %w", err)
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func decodeImage(contentType string, file io.Reader) (image.Image, error) {

	var img image.Image
	var err error
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(file)
	case "image/png":
		img, err = png.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %w", err)
	}

	return img, nil
}

// processImage resizes and compresses the image.
func processImage(img image.Image, contentType string) (*bytes.Buffer, error) {
	maxWidth, maxHeight := uint(800), uint(600)
	img = resize.Thumbnail(maxWidth, maxHeight, img, resize.Lanczos3)

	var buf bytes.Buffer
	switch contentType {
	case "image/jpeg":
		err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 75})
		if err != nil {
			return nil, fmt.Errorf("error compressing JPEG image: %w", err)
		}
	case "image/png":
		err := png.Encode(&buf, img)
		if err != nil {
			return nil, fmt.Errorf("error compressing PNG image: %w", err)
		}
	}

	return &buf, nil
}

func GetProfilePictureUrl(profilePictureKey *string) string {
	if profilePictureKey == nil {
		return ""
	}
	// TODO: Add aws endpoint to config
	url := fmt.Sprintf("http://localhost:4566/%s/%s", "profile-picture", *profilePictureKey)
	return url
}

// UploadProfilePicture handles the HTTP request for uploading a profile picture.
func (r *Resolver) UploadProfilePicture(c *gin.Context) {
	file, _, err := c.Request.FormFile("profilePicture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Could not retrieve the file"})
		c.Abort()
		return
	}
	defer file.Close()

	contentType, err := validateImage(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		c.Abort()
		return
	}

	img, err := decodeImage(contentType, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		c.Abort()
		return
	}

	buf, err := processImage(img, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		c.Abort()
		return
	}

	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Failed to retrieve user email from context"})
		c.Abort()
		return
	}

	// Generate a UUID for the file
	fileUUID := uuid.New().String()
	objectKey := fmt.Sprintf("%s-%s", hashEmail(email.(string)), fileUUID)

	bucket, err := s3utils.GetInstance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		c.Abort()
		return
	}
	if err := bucket.UploadFile(c.Request.Context(), "profile-picture", objectKey, buf, contentType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Failed to upload the file"})
		c.Abort()
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Failed to retrieve user from context"})
		c.Abort()
		return
	}
	userObj, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Invalid user type in context"})
		c.Abort()
		return
	}

	userObj.ProfilePictureKey = &objectKey
	if err := r.UserUsecase.UserRepo.UpdateOne(c, &userObj); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              "success",
		"message":             "Profile picture uploaded successfully",
		"profile_picture_url": fmt.Sprintf("http://localhost:4566/%s/%s", "profile-picture", objectKey),
	})
}
