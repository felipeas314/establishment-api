// config/database.go
package config

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/jackc/pgx/v4"
)

// DB representa a conexão com o banco de dados
var DB *pgx.Conn

// ConnectDB inicializa a conexão com o banco de dados
func ConnectDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", "postgres"),
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_NAME", "backend_db"),
	)

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DB, err = pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	log.Println("✅ Conectado ao PostgreSQL!")
}
