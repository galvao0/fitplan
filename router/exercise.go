package router

import (
	"github.com/galvao0/fitplan/handler/exercise"
	"github.com/gin-gonic/gin"
)

func InitExerciseRoutes(v *gin.RouterGroup) {
	v.GET("/exercise", exercise.ShowExercise)
	v.POST("/exercise", exercise.CreateExercise)
	v.DELETE("/exercise", exercise.DeleteExercise)
	v.PUT("/exercise", exercise.UpdateExercise)
	v.GET("/exercises", exercise.ListExercises)
}
