package controllers

import (
	"net/http"
	"strconv"

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
	var userID uint64 = 0
	categoryID := c.PostForm("category_id")
	title := c.PostForm("title")
	description := c.PostForm("description")

	md := new(models.Quiz)

	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	_, err = md.CreateQuiz(userID, uint64(categoryIDInt), title, description)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	// WIP

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Quiz Success!",
	})
}

func (q QuizController) UpdateQuiz(c *gin.Context) {
	type RequestData struct {
		QuizID uint64 `json:"quiz_id"`
		CategoryID uint64 `json:"quiz_category_id"`
		Title string `json:"title"`
		Description string `json:"description"`
	}

	req := RequestData{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.Quiz)

	err = md.UpdateQuiz(req.QuizID, req.CategoryID, req.Title, req.Description)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}
}

func (q QuizController) DeleteQuiz(c *gin.Context) {
	quizID := c.Param("quiz_id")

	md := new(models.Quiz)

	quizIDInt, err := strconv.Atoi(quizID)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	err = md.DeleteQuiz(uint64(quizIDInt))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}
}
