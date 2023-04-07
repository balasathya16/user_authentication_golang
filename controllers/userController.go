package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	//get email
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
	}
	//hash password

	//create user

	//respond
}
