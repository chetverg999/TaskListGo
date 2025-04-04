package migrations

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	"io/ioutil"
)

func MigrateDatabase(conn *pgxpool.Pool, migrationDir string) {
	files, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		log.Fatalf("Ошибка при чтении директории миграций: %v\n", err)
	}

	for _, file := range files {
		if file.IsDir() || !IsSQLFile(file.Name()) {
			continue
		}
		filePath := fmt.Sprintf("%s/%s", migrationDir, file.Name())
		fmt.Printf("Применяется миграция: %s\n", filePath)
		err := ApplyMigration(conn, filePath)
		if err != nil {
			log.Fatalf("Ошибка при применении миграции %s: %v\n", filePath, err)
		}
	}
	fmt.Println("Миграции применены успешно!")
}

func IsSQLFile(fileName string) bool {

	return len(fileName) > 4 && fileName[len(fileName)-4:] == ".sql"
}

func ApplyMigration(conn *pgxpool.Pool, filePath string) error {
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл миграции %s: %v", filePath, err)
	}
	sql := string(sqlBytes)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("не удалось выполнить миграцию %s: %v", filePath, err)
	}

	return nil
}
