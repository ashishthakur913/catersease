package repositories

import (
	"github.com/ashishthakur913/CaterEase/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (s *UserRepository) CreateUser(user *models.User) error {

	result := s.db.Where(models.User{Email: user.Email}).Assign(&models.User{
		City: user.City,
		Country: user.Country,
		Email: user.Email,
		GoogleUrl: user.GoogleUrl,
		Name: user.Name,
		Notes : user.Notes ,
		PhoneNumber: user.PhoneNumber,
		PostalCode: user.PostalCode,
		State: user.State,
		Street: user.Street,
		Unit: user.Unit ,
	}).FirstOrCreate(&user)

    if result.Error != nil {
        return result.Error
    }

    return nil
}

