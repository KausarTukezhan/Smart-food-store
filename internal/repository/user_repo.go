package repository

import "github.com/KausarTukezhan/Smart-food-store/internal/models"

type UserRepository struct{}

func (r *UserRepository) Create(user models.User) error {
	// TODO: save user
	return nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	return &models.User{}, nil
}
