package api

import (
	"channelMeter/api/controllers"
	"channelMeter/scoreListener"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, scoreRepo *scoreListener.MemoryScoreRepository) {
	student := controllers.StudentController{ScoreRepo: scoreRepo}
	exam := controllers.ExamController{ScoreRepo: scoreRepo}

	router.GET("/students", student.GetStudents)
	router.GET("/students/:id", student.GetStudent)
	router.GET("/exams", exam.GetExams)
	router.GET("/exams/:id", exam.GetExam)
}
