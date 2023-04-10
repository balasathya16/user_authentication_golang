package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("in middleware")
}
