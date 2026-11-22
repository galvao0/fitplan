package workoutexercise

import (
	"fmt"
	"net/http"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func AddExerciseWorkout(c *gin.Context) {
	workoutID := c.Param("workoutID")
	if workoutID == "" {
		handler.SendError(c, http.StatusBadRequest, handler.ErrParamIsRequired("workoutID", "queryParam").Error())
		return
	}

	exerciseID := c.Param("exerciseID")
	if exerciseID == "" {
		handler.SendError(c, http.StatusBadRequest, handler.ErrParamIsRequired("exerciseID", "queryParam").Error())
		return
	}

	workout := schemas.Workout{}
	if err := handler.GetDB().First(&workout, workoutID).Error; err != nil {
		handler.SendError(c, http.StatusNotFound, fmt.Sprintf("workout with id: %s not found", workoutID))
		return
	}

	exercise := schemas.Exercise{}
	if err := handler.GetDB().First(&exercise, exerciseID).Error; err != nil {
		handler.SendError(c, http.StatusNotFound, fmt.Sprintf("exercise with id: %s not found", exerciseID))
		return
	}

	if err := handler.GetDB().Model(&workout).Association("Exercises").Append(&exercise); err != nil {
		handler.SendError(c, http.StatusInternalServerError, "error association workout with exercise")
		return
	}

	handler.SendSuccess(c, "add-exercise-workout", workout)
}
