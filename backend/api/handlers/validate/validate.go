package validate

import (
	"net/http"

	"github.com/SantGT5/quintosgo/internal/models"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	user, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	userData, ok := user.(models.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Im logged in",
		"user": gin.H{
			"email":     userData.Email,
			"CreatedAt": userData.CreatedAt,
		},
	})
}
