package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (t TestController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
