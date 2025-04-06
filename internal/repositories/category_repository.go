package repositories

import (
	"context"
	"establishment/config"
	"establishment/internal/models"
	"log"
)

func CreateCategory(category models.Category) error {
	query := `INSERT INTO category (name, description, establishment_id) values ($1, $2, $3)`
	_, err := config.DB.Exec(context.Background(), query, category.Name, category.Description, category.EstablishmentID)
	if err != nil {
		log.Println("Erro ao criar categoria do restaurante:", err)
		return err
	}
	return nil
}

func GetGategoriesByEstablishment(page, limit int) ([]models.Category, int, error) {
	offset := (page - 1) * limit

	query := `SELECT id, name, description FROM category ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := config.DB.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	categories := []models.Category{}
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.EstablishmentID)
		if err != nil {
			return nil, 0, err
		}
		categories = append(categories, category)
	}

	var total int
	err = config.DB.QueryRow(context.Background(), `SELECT COUNT(*) FROM category`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}
