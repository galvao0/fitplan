package exercise

import (
	"fmt"

	"github.com/galvao0/fitplan/handler"
)

type ExerciseRequest struct {
	Name   string `json:"name"`
	Weight uint64 `json:"weight"`
	Reps   uint16 `json:"reps"`
	Sets   uint16 `json:"sets"`
}

type CreateExerciseRequest struct {
	ExerciseRequest
}

type UpdateExerciseRequest struct {
	ExerciseRequest
}

func (r *CreateExerciseRequest) Validate() error {
	if r.Name == "" {
		return handler.ErrParamIsRequired("name", "string")
	}
	if r.Weight == 0 {
		return handler.ErrParamIsRequired("weight", "uint64")
	}
	if r.Reps == 0 {
		return handler.ErrParamIsRequired("reps", "uint16")
	}
	if r.Sets == 0 {
		return handler.ErrParamIsRequired("sets", "uint16")
	}
	return nil
}

func (r *UpdateExerciseRequest) Validate() error {
	if r.Name != "" || r.Weight != 0 || r.Reps != 0 || r.Sets != 0 {
		return nil
	}
	return fmt.Errorf("at last one valid field must be provided")
}
