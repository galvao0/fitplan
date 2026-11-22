package router

import (
	"github.com/galvao0/fitplan/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	handler.Init()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/greeting", handler.Greeting)
		InitWorkoutRoutes(v1)
	}
}
