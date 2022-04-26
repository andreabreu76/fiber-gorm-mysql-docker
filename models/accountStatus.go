package models

import "gorm.io/gorm"

type AccountStatus struct {
	gorm.Model
	Name string `json:"name"`
}
