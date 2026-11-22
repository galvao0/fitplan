package workout

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateWorkout(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		handler.SendError(c, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParam").Error())
		return
	}

	request := UpdateWorkoutRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	workout := schemas.Workout{}

	if err := handler.GetDB().Preload("Exercises").First(&workout, id).Error; err != nil {
		handler.SendError(c, http.StatusNotFound, fmt.Sprintf("workout with id: %s not found", id))
		return
	}

	normalize := strings.ToLower(request.Name)

	if err := handler.GetDB().Where("name = ?", normalize).First(&schemas.Workout{}).Error; err == nil {
		handler.SendError(c, http.StatusConflict, fmt.Sprintf("workout with name: %s already exists", normalize))
		return
	}

	if request.Name != "" {
		workout.Name = normalize
	}
	if request.Image != "" {
		workout.Image = request.Image
	}

	if err := handler.GetDB().Save(&workout).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccess(c, "update-workout", workout)
}
