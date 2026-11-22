package exercise

import (
	"net/http"
	"strings"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func CreateExercise(c *gin.Context) {
	request := CreateExerciseRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	exNameFormat := strings.ToLower(request.Name)

	if err := handler.GetDB().Where("name = ?", exNameFormat).First(&schemas.Exercise{}).Error; err == nil {
		handler.SendError(c, http.StatusConflict, "exercise already exists")
		return
	}

	exercise := schemas.Exercise{
		Name:   exNameFormat,
		Weight: request.Weight,
		Reps:   request.Reps,
		Sets:   request.Sets,
	}

	if err := handler.GetDB().Create(&exercise).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccess(c, "create-exercise", exercise)
}
