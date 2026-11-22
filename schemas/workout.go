package schemas

import "gorm.io/gorm"

type Workout struct {
	gorm.Model
	Name      string
	Image     string
	Exercises []Exercise `gorm:"many2many:workout_exercises"`
}
