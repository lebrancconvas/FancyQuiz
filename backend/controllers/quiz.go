package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/FancyQuiz/models"
	"github.com/lebrancconvas/FancyQuiz/utils"
)

type QuizController struct{}

func (q QuizController) GetAllQuizHeader(c *gin.Context) {
	md := new(models.Quiz)

	res, err := md.GetAllQuizHeader()
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Quiz Header Success!",
		"data": res,
	})
}

func (q QuizController) GetAllQuiz(c *gin.Context) {
	md := new(models.Quiz)

	res, err := md.GetAllQuiz()
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Quiz Success!",
		"data": res,
	})
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
		"message": "Get All Quiz Category Success!",
		"data": res,
	})
}

func (q QuizController) CreateQuiz(c *gin.Context) {
	var userID uint64 = 0
	categoryID, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}


	title := c.PostForm("title")
	description := c.PostForm("description")

	md := new(models.Quiz)
	_, err = md.CreateQuiz(userID, uint64(categoryID), title, description)
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
		CategoryID uint64 `json:"quiz_category_id"`
		Title string `json:"title"`
		Description string `json:"description"`
	}

	quizID, err := strconv.Atoi(c.Param("quiz_id"))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	req := RequestData{}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.Quiz)

	err = md.UpdateQuiz(uint64(quizID), req.CategoryID, req.Title, req.Description)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Quiz Success!",
	})
}

func (q QuizController) DeleteQuiz(c *gin.Context) {
	quizID, err := strconv.Atoi(c.Param("quiz_id"))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.Quiz)

	err = md.DeleteQuiz(uint64(quizID))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Quiz Success!",
	})
}
