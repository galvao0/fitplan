package workout

import (
	"net/http"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func ListWorkouts(c *gin.Context) {
	workouts := []schemas.Workout{}

	if err := handler.GetDB().Find(&workouts).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, "error listing workouts")
		return
	}

	handler.SendSuccess(c, "list-workouts", workouts)
}
