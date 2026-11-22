package schemas

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name   string
	Weight uint64
	Reps   uint16
	Sets   uint16
}
