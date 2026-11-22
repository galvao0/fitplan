package handler

import (
	"fmt"
	"net/http"

	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteWorkout(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	workout := schemas.Workout{}

	if err := db.First(&workout, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("workout with id: %s not found", id))
		return
	}

	if err := db.Delete(&workout).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting workout with id: %s", id))
		return
	}

	sendSuccess(c, "delete-workout", workout)
}
