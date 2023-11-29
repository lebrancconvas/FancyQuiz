package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/lebrancconvas/FancyQuiz/forms"
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
		"message": "Get All Users Success!",
		"data": res,
	})
}

func (u UserController) GetUserInformation(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.User)
	res, err := md.GetUserInformation(uint64(userID))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get User Information Success!",
		"data": res,
	})
}

func (u UserController) CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	// password := c.PostForm("password")
	displayName := c.PostForm("display_name")

	md := new(models.User)
	res, err := md.CreateUser(username, displayName)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create User Success!",
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Update User Success!",
	})
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User Success!",
	})
}
