package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	// get the cookie from request

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode & validate

	// check the exp

	// Find the user with token in subject

	// Attach token to request
	//continue

	c.Next()
}
