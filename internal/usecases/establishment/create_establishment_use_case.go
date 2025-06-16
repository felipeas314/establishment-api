package usecasesestablishment

import (
	"establishment/internal/models"
	"establishment/internal/repositories"
)

func CreateEstablishment(establishment models.Establishment) error {
	return repositories.CreateEstablishment(establishment)
}
