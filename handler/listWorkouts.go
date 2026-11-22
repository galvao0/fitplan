package handler

import (
	"net/http"

	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func ListWorkouts(c *gin.Context) {
	workouts := []schemas.Workout{}

	if err := db.Find(&workouts).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing workouts")
		return
	}

	sendSuccess(c, "list-workouts", workouts)
}
