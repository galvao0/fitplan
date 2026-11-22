package handler

import (
	"fmt"
	"net/http"

	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func ShowWorkout(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	workout := schemas.Workout{}

	if err := db.First(&workout, id).Error; err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("workout with id: %s not found", id))
		return
	}

	sendSuccess(c, "show-workout", workout)
}
