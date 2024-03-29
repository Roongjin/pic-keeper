package user

import (
	"net/http"

	"github.com/Roongkun/software-eng-ii/internal/controller/util"
	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/Roongkun/software-eng-ii/internal/third-party/auth"
	"github.com/Roongkun/software-eng-ii/internal/third-party/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Resolver) GoogleCallback(c *gin.Context) {
	provider := "Google"
	config := util.GetGoogleLibConfig(c)
	code := c.Query("code")

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Authorization code not provided",
		})
		c.Abort()
		return
	}

	tokenRes, err := oauth2.GetGoogleOAuth2Token(config, code)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	googleUser, err := oauth2.GetGoogleUser(tokenRes.AccessToken, tokenRes.IdToken)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	exist, err := r.UserUsecase.CheckExistenceByEmail(c, googleUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	var user *model.User
	if !exist {
		newUser := model.User{
			Id:        uuid.New(),
			Name:      googleUser.Name,
			Email:     googleUser.Email,
			Provider:  &provider,
			Password:  nil,
			LoggedOut: false,
		}

		if err := r.UserUsecase.UserRepo.AddOne(c, &newUser); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		user = &newUser
	} else {
		var innerErr error
		user, innerErr = r.UserUsecase.UserRepo.FindOneByEmail(c, googleUser.Email)
		if innerErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		user.LoggedOut = false
		if err := r.UserUsecase.UserRepo.UpdateOne(c, user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}

	secretKey, exist := c.Get("secretKey")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "secret key not found",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         secretKey.(string),
		Issuer:            "AuthProvider",
		ExpirationMinutes: 5,
		ExpirationHours:   12,
	}

	token, err := jwtWrapper.GenerateToken(c, googleUser.Email, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"session_token": token,
	})
}
