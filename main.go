package main

import (
	"net/http"
	initialize "userauth/init"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnvVariables()
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
