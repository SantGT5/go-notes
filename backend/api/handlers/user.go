package handlers

import (
	"net/http"

	repo "github.com/SantGT5/quintosgo/internal/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostUser(c *gin.Context) {
	var err error
	var hash []byte

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err = bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := repo.User{Email: body.Email, Password: string(hash)}

	err = user.CreateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": gin.H{
				"email": user.Email,
			},
		},
	})
}
