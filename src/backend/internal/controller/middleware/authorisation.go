package middleware

import (
	"net/http"
	"strings"

	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/Roongkun/software-eng-ii/internal/third-party/auth"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	var token string
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		// If the Authorization header is not present, return a 403 status code
		c.JSON(http.StatusForbidden, gin.H{"error": "No Authorization header provided"})
		c.Abort()
		return
	}
	// Split the Authorization header to get the token
	extractedToken := strings.Split(authorizationHeader, "Bearer ")
	if len(extractedToken) == 2 {
		// Trim the token
		token = strings.TrimSpace(extractedToken[1])
	} else {
		// If the token is not in the correct format, return a 400 status code
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Format of Authorization Token"})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey: c.Request.Context().Value(model.ContextKey("secretKey")).(string),
		Issuer:    "AuthProvider",
	}

	claims, err := jwtWrapper.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Next()
}