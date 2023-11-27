package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/FancyQuiz/models"
	"github.com/lebrancconvas/FancyQuiz/utils"
)

type ReportController struct{}

func (r ReportController) GetAllReport(c *gin.Context) {
	md := new(models.Report)

	res, err := md.GetAllReport()
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (r ReportController) GetReportFromDateCreated(c *gin.Context) {

}

func (r ReportController) CreateReport(c *gin.Context) {

}
