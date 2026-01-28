package repository

import "github.com/KausarTukezhan/Smart-food-store/internal/models"

type ProductRepository struct{}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	return &models.Product{}, nil
}
