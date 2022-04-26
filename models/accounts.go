package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	BankName string        `json:"bank_name" validate:"required"`
	Number   int           `json:"number" validate:"required"`
	Digit    int64         `json:"digits" validate:"required"`
	IDStatus int           `json:"id_status"`
	Status   AccountStatus `json:"status" gorm:"foreignkey:IDStatus"`
}
