// internal/repositories/user_repository.go
package repositories

import (
	"context"
	"log"

	"establishment/config"
	"establishment/internal/models"
)

// Criar usuário no banco
func CreateEstablishment(establishment models.Establishment) error {
	query := `INSERT INTO establishment (name, description) VALUES ($1, $2)`

	_, err := config.DB.Exec(context.Background(), query, establishment.Name, establishment.Description)
	if err != nil {
		log.Println("Erro ao criar usuário:", err)
		return err
	}
	return nil
}
