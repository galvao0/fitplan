package handler

import (
	"net/http"

	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateWorkout(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	request := UpdateWorkoutRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	workout := schemas.Workout{}

	if err := db.First(&workout, id).Error; err != nil {
		sendError(c, http.StatusBadRequest, "workout not found")
		return
	}

	if request.Name != "" {
		workout.Name = request.Name
	}
	if request.Image != "" {
		workout.Image = request.Image
	}

	if err := db.Save(&workout).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "update-workout", workout)
}
