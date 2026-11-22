package router

import (
	"github.com/galvao0/fitplan/handler/workout"
	"github.com/gin-gonic/gin"
)

func InitWorkoutRoutes(v *gin.RouterGroup) {
	v.GET("/workout", workout.ShowWorkout)
	v.POST("/workout", workout.CreatingWorkout)
	v.DELETE("/workout", workout.DeleteWorkout)
	v.PUT("/workout", workout.UpdateWorkout)
	v.GET("/workouts", workout.ListWorkouts)
}
