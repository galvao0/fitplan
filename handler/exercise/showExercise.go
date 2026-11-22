package exercise

import (
	"fmt"
	"net/http"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func ShowExercise(c *gin.Context) {
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

	handler.SendSuccess(c, "show-exercise", exercise)
}
