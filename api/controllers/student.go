package controllers

import (
	"channelMeter/scoreListener"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	ScoreRepo scoreListener.ScoreRepository
}

func (ctrl StudentController) GetStudents(c *gin.Context) {
	c.JSON(200, ctrl.ScoreRepo.GetStudents())
}

func (ctrl StudentController) GetStudent(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, ctrl.ScoreRepo.GetStudentScores(id))
}
