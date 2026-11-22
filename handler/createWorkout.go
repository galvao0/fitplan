package handler

import (
	"net/http"

	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func CreatingWorkout(c *gin.Context) {
	request := CreateWorkoutRequest{}
	c.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Where("name = ?", request.Name).First((&schemas.Workout{})).Error; err == nil {
		sendError(c, http.StatusConflict, "workout already exists")
		return
	}

	workout := schemas.Workout{
		Name:  request.Name,
		Image: request.Image,
	}

	if err := db.Create(&workout).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "create-workout", workout)
}
