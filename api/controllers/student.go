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
	student := ctrl.ScoreRepo.GetStudentScores(id)

	if len(student.StudentId) == 0 {
		c.AbortWithStatus(400)
	}

	c.JSON(200, student)
}
