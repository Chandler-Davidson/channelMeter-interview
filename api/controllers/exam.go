package controllers

import (
	"channelMeter/scoreListener"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamController struct {
	ScoreRepo scoreListener.ScoreRepository
}

func (ctrl ExamController) GetExams(c *gin.Context) {
	c.JSON(200, ctrl.ScoreRepo.GetExams())
}

func (ctrl ExamController) GetExam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, ctrl.ScoreRepo.GetExam(id))
}
