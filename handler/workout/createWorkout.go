package workout

import (
	"net/http"
	"strings"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func CreatingWorkout(c *gin.Context) {
	request := CreateWorkoutRequest{}
	c.BindJSON(&request)
	if err := request.Validate(); err != nil {
		handler.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	nameFormat := strings.ToLower(request.Name)

	if err := handler.GetDB().Where("name = ?", nameFormat).First((&schemas.Workout{})).Error; err == nil {
		handler.SendError(c, http.StatusConflict, "workout already exists")
		return
	}

	workout := schemas.Workout{
		Name:  nameFormat,
		Image: request.Image,
	}

	if err := handler.GetDB().Create(&workout).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccess(c, "create-workout", workout)
}
