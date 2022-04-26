package models

import "gorm.io/gorm"

type Moviments struct {
	gorm.Model
	IDAccount int          `json:"id_account"`
	Account   Account      `json:"account" gorm:"foreignkey:IDAccount"`
	IDType    int          `json:"id_type"`
	Type      MovimentType `json:"type" gorm:"foreignkey:IDType"`
	Value     float64      `json:"value" gorm:"type:decimal(10,2)"`
}
