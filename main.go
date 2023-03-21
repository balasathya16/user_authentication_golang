package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user_authentication_golang/initializers"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
