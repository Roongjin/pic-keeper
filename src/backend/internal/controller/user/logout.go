package user

import (
	"net/http"

	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/gin-gonic/gin"
)

func (r *Resolver) Logout(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"error":  "the user is no longer existed",
		})
		c.Abort()
		return
	}

	userModel, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "please check the model of the user",
			"error":   "cannot do type assertion",
		})
		c.Abort()
		return
	}

	userModel.LoggedOut = true
	if err := r.UserUsecase.UserRepo.UpdateOne(c, &userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "you have logged out successfully",
	})
}
