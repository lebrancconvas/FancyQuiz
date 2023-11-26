package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/FancyQuiz/models"
	"github.com/lebrancconvas/FancyQuiz/utils"
)

type UserController struct{}

func (u UserController) GetAllUsers(c *gin.Context) {
	md := new(models.User)
	res, err := md.GetAllUsers()
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
