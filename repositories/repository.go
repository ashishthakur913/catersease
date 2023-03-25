package repositories

import (
	"gorm.io/gorm"

	"github.com/ashishthakur913/CaterEase/models"
)

type DBServer struct {
	db *gorm.DB
}

func NewDBServer(db *gorm.DB) (*DBServer, error) {
	// Migrate the schema
	err := db.AutoMigrate(&models.User{}, &models.Order{}, &models.Driver{})
	if err != nil {
		return nil, err
	}

	return &DBServer{db}, nil
}
