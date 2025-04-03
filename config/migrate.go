package config

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	dsn := fmt.Sprintf(
		"pgx://%s:%s@%s:%s/%s?sslmode=disable",
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", "postgres"),
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_NAME", "backend_db"),
	)

	m, err := migrate.New(
		"file://migrations", // Caminho relativo ao projeto
		dsn,
	)
	if err != nil {
		log.Fatalf("Erro ao criar o migrator: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao aplicar migrations: %v", err)
	}

	log.Println("✅ Migrations aplicadas com sucesso!")
}
