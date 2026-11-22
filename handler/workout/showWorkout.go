package workout

import (
	"fmt"
	"net/http"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func ShowWorkout(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		handler.SendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	workout := schemas.Workout{}

	if err := handler.GetDB().First(&workout, id).Error; err != nil {
		handler.SendError(c, http.StatusBadRequest, fmt.Sprintf("workout with id: %s not found", id))
		return
	}

	handler.SendSuccess(c, "show-workout", workout)
}
