package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"userauth/initializers"
	"userauth/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	// get the cookie from request

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode, validate

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {


		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

	})


		//check expiration

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}
		// find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

		//attach to request

		c.Set("user", user)

		//continue

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

}
