package workout

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type WorkoutRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type CreateWorkoutRequest struct{ WorkoutRequest }
type UpdateWorkoutRequest struct{ WorkoutRequest }

func (r *WorkoutRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Image == "" {
		r.Image = "not"
	}
	return nil
}
