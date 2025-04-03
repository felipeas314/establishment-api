// internal/services/user_service.go
package usecases

import (
	"establishment/internal/models"
	"establishment/internal/repositories"
)

// Criar um usuário chamando o repositório
func CreateEstablishment(establishment models.Establishment) error {
	return repositories.CreateEstablishment(establishment)
}
