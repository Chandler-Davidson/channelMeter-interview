package main

import (
	"channelMeter/api"
	"channelMeter/scoreListener"

	"github.com/gin-gonic/gin"
)

func main() {
	scoreRepo := scoreListener.NewScoreRepository()
	router := gin.Default()

	api.RegisterRoutes(router, &scoreRepo)
	router.Run("localhost:8080")
}
