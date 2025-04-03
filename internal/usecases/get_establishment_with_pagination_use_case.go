package usecases

import (
	"establishment/internal/models"
	"establishment/internal/repositories"
)

func GetEstablishmentWithPagination(page, limit int) ([]models.Establishment, int, error) {
	return repositories.GetUsersPaginated(page, limit)
}
