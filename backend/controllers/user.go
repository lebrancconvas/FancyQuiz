package controllers

import (
	"net/http"
	"strconv"

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
	res, err := md.CreateUser(userInformation)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (u UserController) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	type RequestData struct {
		DisplayName 	string `json:"display_name"`
	}

	req := RequestData{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.User)

	err = md.UpdateUser(uint64(userID), req.DisplayName)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}
}

func (u UserController) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.User)

	err = md.DeleteUser(uint64(userID))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}
}
