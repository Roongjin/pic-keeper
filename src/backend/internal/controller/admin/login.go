package admin

import (
	"net/http"

	"github.com/Roongkun/software-eng-ii/internal/controller/admin/fieldvalidate"
	"github.com/Roongkun/software-eng-ii/internal/controller/util"
	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/Roongkun/software-eng-ii/internal/third-party/auth"
	"github.com/gin-gonic/gin"
)

func (r *Resolver) Login(c *gin.Context) {
	cred := model.LoginCredentials{}
	if err := c.BindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"error":   err.Error(),
			"message": "unable to bind request body with json, please recheck",
		})
		c.Abort()
		return
	}

	if fieldErr := fieldvalidate.LoginAdmin(cred); len(fieldErr) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"errors": util.JSONErrs(fieldErr),
		})
		c.Abort()
		return
	}

	existedAdmin, err := r.AdminUsecase.FindOneByEmail(c, cred.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}

	secretKey, exist := c.Get("adminSecretKey")
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

	existedAdmin.LoggedOut = false
	if err := r.AdminUsecase.AdminRepo.UpdateOne(c, existedAdmin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}

	token, err := jwtWrapper.GenerateToken(c, cred.Email, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"session-token": token,
	})

}
