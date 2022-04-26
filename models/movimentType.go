package models

import "gorm.io/gorm"

type MovimentType struct {
	gorm.Model
	Name string `json:"name"`
}
