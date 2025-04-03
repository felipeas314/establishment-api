// internal/repositories/user_repository.go
package repositories

import (
	"context"
	"log"

	"establishment/config"
	"establishment/internal/models"
)

func CreateEstablishment(establishment models.Establishment) error {
	query := `INSERT INTO establishment (name, description) VALUES ($1, $2)`

	_, err := config.DB.Exec(context.Background(), query, establishment.Name, establishment.Description)
	if err != nil {
		log.Println("Erro ao criar usuário:", err)
		return err
	}
	return nil
}

func GetUsersPaginated(page, limit int) ([]models.Establishment, int, error) {
	offset := (page - 1) * limit

	query := `SELECT id, name, description FROM establishment ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := config.DB.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	establishments := []models.Establishment{}
	for rows.Next() {
		var establishment models.Establishment
		err := rows.Scan(&establishment.ID, &establishment.Name, &establishment.Description)
		if err != nil {
			return nil, 0, err
		}
		establishments = append(establishments, establishment)
	}

	var total int
	err = config.DB.QueryRow(context.Background(), `SELECT COUNT(*) FROM establishment`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return establishments, total, nil
}
