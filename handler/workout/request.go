package workout

import "github.com/galvao0/fitplan/handler"

type WorkoutRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type CreateWorkoutRequest struct{ WorkoutRequest }
type UpdateWorkoutRequest struct{ WorkoutRequest }

func (r *WorkoutRequest) Validate() error {
	if r.Name == "" {
		return handler.ErrParamIsRequired("name", "string")
	}
	if r.Image == "" {
		r.Image = "not"
	}
	return nil
}
