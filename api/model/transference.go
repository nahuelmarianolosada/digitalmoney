package model

import (
	"gorm.io/gorm"
	"time"
)

type Transference struct {
	gorm.Model
	AccountID   uint
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Amount      int       `json:"ammount"`
	Dated       time.Time `json:"dated"`
}

type TransferenceRequest struct {
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Amount      int       `json:"ammount"`
	Dated       time.Time `json:"dated"`
}
