package usecasescategory

import (
	"establishment/internal/models"
	"establishment/internal/repositories"
)

func CreateCategory(category models.Category) error {
	return repositories.CreateCategory(category)
}
