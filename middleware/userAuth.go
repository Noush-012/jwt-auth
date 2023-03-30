package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Noush-012/JWT/initializer"
	"github.com/Noush-012/JWT/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func UserAuth(c *gin.Context) {
	fmt.Println("In Middleware")
	// Get the cookie of request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Validation user

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWTCODE")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if expired or not
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find the user with token sub
		var user models.User
		initializer.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Attach to request
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
