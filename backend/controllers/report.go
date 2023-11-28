package controllers

import (
	"net/http"
	"strconv"

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
	type RequestData struct {
		UserID uint64 `json:"userID" binding:"required"`
		ReportContent string `json:"content"`
	}

	req := RequestData{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.Report)

	err = md.CreateReport(req.UserID, req.ReportContent)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Report created successfully",
	})
}

func (r ReportController) DeleteReport(c *gin.Context) {
	reportID := c.Param("report_id")

	reportIDInt, err := strconv.Atoi(reportID)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.Report)

	err = md.DeleteReport(uint64(reportIDInt))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}
}

func (r ReportController) AcceptReport(c *gin.Context) {
	reportID := c.Param("report_id")

	reportIDInt, err := strconv.Atoi(reportID)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.Report)

	err = md.UpdateReportToBeAccepted(uint64(reportIDInt))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}
}

func (r ReportController) CompleteReport(c *gin.Context) {
	reportID := c.Param("report_id")

	reportIDInt, err := strconv.Atoi(reportID)
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}

	md := new(models.Report)

	err = md.UpdateReportToBeCompleted(uint64(reportIDInt))
	if err != nil {
		utils.UnprocessableLog(c, err)
		return
	}
}


