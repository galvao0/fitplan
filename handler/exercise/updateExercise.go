package exercise

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/galvao0/fitplan/handler"
	"github.com/galvao0/fitplan/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateExercise(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		handler.SendError(c, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParam").Error())
		return
	}

	request := UpdateExerciseRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.SendError(c, http.StatusBadRequest, err.Error())
	}

	nameFormat := strings.ToLower(request.Name)

	if err := handler.GetDB().Where("name = ?", nameFormat).First(&schemas.Exercise{}).Error; err == nil {
		handler.SendError(c, http.StatusConflict, fmt.Sprintf("exercise with name: %s already exists", nameFormat))
		return
	}

	exercise := schemas.Exercise{}

	if err := handler.GetDB().First(&exercise, id).Error; err != nil {
		handler.SendError(c, http.StatusNotFound, fmt.Sprintf("exercise with id: %s not found", id))
		return
	}

	if request.Name != "" {
		exercise.Name = nameFormat
	}
	if request.Weight != 0 {
		exercise.Weight = request.Weight
	}
	if request.Reps != 0 {
		exercise.Reps = request.Reps
	}
	if request.Sets != 0 {
		exercise.Sets = request.Sets
	}

	if err := handler.GetDB().Save(&exercise).Error; err != nil {
		handler.SendError(c, http.StatusInternalServerError, "error updating exercise")
		return
	}

	handler.SendSuccess(c, "update-exercise", exercise)
}
