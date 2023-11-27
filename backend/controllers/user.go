package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/FancyQuiz/forms"
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

func (u UserController) CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	displayName := c.PostForm("display_name")

	var userInformation forms.UserRegister = forms.UserRegister{
		Username: username,
		Password: password,
		DisplayName: displayName,
	}

	md := new(models.User)
	err := md.CreateUser(userInformation)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create User Success!",
	})
}
