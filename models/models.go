package models

import (
  "gorm.io/gorm"
)

type User struct {
	gorm.Model
	City string `json:"city"`
	Country string `json:"country"`
	Email string `json:"email" gorm:"unique"`
	GoogleUrl string `json:"google_url"`
	Name string `json:"name"`
	Notes *string `json:"notes,omitempty"`
	PhoneNumber string `json:"phone_number"`
	PostalCode string `json:"postal_code"`
	State string `json:"state"`
	Street string `json:"street"`
	Unit *string `json:"unit,omitempty"`
}

type Order struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	DriverID uint   `gorm:"not null"`
	Status string `gorm:"not null"`
	Notes  string
}

type Driver struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phone_number"`
}
