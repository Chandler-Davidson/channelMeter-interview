package main

import (
	"channelMeter/scores/scores"
	"strconv"

	"github.com/gin-gonic/gin"
)

var scoreRepo scores.ScoreDataStore

func main() {
	scoreRepo = scores.NewDataStore()
	router := gin.Default()

	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudent)
	router.GET("/exams", getExams)
	router.GET("/exams/:id", getExam)

	router.Run("localhost:8080")
}

func getStudents(c *gin.Context) {
	c.JSON(200, scoreRepo.GetStudents())
}

func getStudent(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, scoreRepo.GetStudentScores(id))
}

func getExams(c *gin.Context) {
	c.JSON(200, scoreRepo.GetExams())
}

func getExam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, nil)
		return
	}

	c.JSON(200, scoreRepo.GetExam(id))
}
