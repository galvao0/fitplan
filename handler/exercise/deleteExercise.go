package exercise

import (
	"fmt"
	"net/http"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteExercise(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		handler.SendError(c, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParam").Error())
		return
	}

	exercise := schemas.Exercise{}

	if err := handler.GetDB().First(&exercise, id).Error; err != nil {
		handler.SendError(c, http.StatusNotFound, fmt.Sprintf("exercise with id: %s not found", id))
		return
	}

	if err := handler.GetDB().Delete(&exercise).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting with id: %s", id))
		return
	}

	handler.SendSuccess(c, "delete-exercise", exercise)
}
