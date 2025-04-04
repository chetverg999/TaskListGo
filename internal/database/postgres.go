package database

import (
	"awesomeProject2/migrations"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *pgxpool.Pool

func ConnectDB() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	migrationDir := "./migrations"
	migrations.MigrateDatabase(dbpool, migrationDir)
	DB = dbpool
	fmt.Println("Подключение к БД успешно")
}
