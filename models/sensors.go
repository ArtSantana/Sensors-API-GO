package models

import (
	"github.com/jinzhu/gorm"
)

// Sensor model
type Sensor struct {
	gorm.Model
	Temperature    float64
	Active         bool
	MaxTemperature float64
}

// UpdateModel model
type UpdateModel struct {
	ID          int
	Temperature float64
}
