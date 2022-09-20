package model

import (
	"gorm.io/gorm"
	"time"
)

type Card struct {
	gorm.Model
	AccountID      uint
	NumberID       int       `json:"number_id"`
	FirstLastname  string    `json:"first_last_name"`
	Cod            int       `json:"cod"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type CardRequest struct {
	NumberID       int       `json:"number_id"`
	FirstLastname  string    `json:"first_last_name"`
	Cod            int       `json:"cod"`
	ExpirationDate time.Time `json:"expiration_date"`
}
