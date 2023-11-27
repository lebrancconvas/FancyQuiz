package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/FancyQuiz/models"
	"github.com/lebrancconvas/FancyQuiz/utils"
)

type QuizController struct{}

func (q QuizController) GetAllQuiz(c *gin.Context) {

}

func (q QuizController) GetAllQuizFromCreatedUser(c *gin.Context) {

}

func (q QuizController) GetAllQuizFromParticipatedUser(c *gin.Context) {

}

func (q QuizController) GetAllQuizCategory(c *gin.Context) {
	md := new(models.Quiz)
	res, err := md.GetAllQuizCategory()
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (q QuizController) CreateQuiz(c *gin.Context) {

}

func (q QuizController) UpdateQuiz(c *gin.Context) {

}

func (q QuizController) DeleteQuiz(c *gin.Context) {

}
