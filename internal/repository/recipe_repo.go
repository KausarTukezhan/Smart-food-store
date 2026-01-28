package repository

import "github.com/KausarTukezhan/Smart-food-store/internal/models"

type RecipeRepository struct{}

func (r *RecipeRepository) GetByID(id int) (*models.Recipe, error) {
	return &models.Recipe{}, nil
}
