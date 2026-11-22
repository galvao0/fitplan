package workout

import (
	"fmt"
	"net/http"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteWorkout(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		handler.SendError(c, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParam").Error())
		return
	}

	workout := schemas.Workout{}

	if err := handler.GetDB().First(&workout, id).Error; err != nil {
		handler.SendError(c, http.StatusNotFound, fmt.Sprintf("workout with id: %s not found", id))
		return
	}

	if err := handler.GetDB().Delete(&workout).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting workout with id: %s", id))
		return
	}

	handler.SendSuccess(c, "delete-workout", workout)
}
