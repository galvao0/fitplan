package router

import (
	"github.com/galvao0/fitplan/handler"
	"github.com/gin-gonic/gin"
)

func InitWorkoutRoutes(v *gin.RouterGroup) {
	v.GET("/workout", handler.ShowWorkout)
	v.POST("/workout", handler.CreatingWorkout)
	v.DELETE("/workout", handler.DeleteWorkout)
	v.PUT("/workout", handler.UpdateWorkout)
	v.GET("/workouts", handler.ListWorkouts)
}
