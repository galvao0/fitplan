package exercise

import (
	"net/http"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func ListExercises(c *gin.Context) {
	exercises := []schemas.Exercise{}

	if err := handler.GetDB().Find(&exercises).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, "error listing exercises")
		return
	}

	handler.SendSuccess(c, "list-exercises", exercises)
}
