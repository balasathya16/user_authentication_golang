package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	// get the cookie from request

	//decode and validate cookie

	//check expiration

	// find the user with token

	// attach to request

	// continue

	c.Next()
}
