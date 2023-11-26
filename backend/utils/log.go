package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnprocessableLog(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{})
	log.Fatal(err.Error())
}
