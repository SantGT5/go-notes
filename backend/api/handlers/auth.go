package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/SantGT5/quintosgo/internal/database"
	"github.com/SantGT5/quintosgo/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to read body",
		})
		return
	}

	var user models.User
	db := database.GetDatabase()

	db.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, false)

	c.JSON(http.StatusOK, gin.H{"data": gin.H{
		"user": gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	}})
}
