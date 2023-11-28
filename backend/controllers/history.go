package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/FancyQuiz/models"
	"github.com/lebrancconvas/FancyQuiz/utils"
)

type HistoryController struct{}

func (h HistoryController) GetAllHistory(c *gin.Context) {

}

func (h HistoryController) GetHistoryFromUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.History)

	res, err := md.GetAllHistoryFromUser(uint64(userID))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})

}

func (h HistoryController) CreateHistory(c *gin.Context) {

}
