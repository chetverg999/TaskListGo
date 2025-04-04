package migrations

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	pgxpool2 "github.com/jackc/pgx/v5/pgxpool"
	"io/ioutil"
)

// Функция для загрузки и применения миграции
func MigrateDatabase(conn *pgxpool2.Pool, migrationDir string) {
	// Чтение файлов миграций
	files, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		log.Fatalf("Ошибка при чтении директории миграций: %v\n", err)
	}

	// Применение миграций
	for _, file := range files {
		if file.IsDir() || !IsSQLFile(file.Name()) {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", migrationDir, file.Name())
		fmt.Printf("Применяется миграция: %s\n", filePath)

		// Применяем миграцию
		err := ApplyMigration(conn, filePath)
		if err != nil {
			log.Fatalf("Ошибка при применении миграции %s: %v\n", filePath, err)
		}
	}

	fmt.Println("Миграции применены успешно!")
}

// Проверка, что файл является SQL миграцией
func IsSQLFile(fileName string) bool {
	return len(fileName) > 4 && fileName[len(fileName)-4:] == ".sql"
}

// Применение SQL миграции
func ApplyMigration(conn *pgxpool2.Pool, filePath string) error {
	// Чтение SQL запроса из файла
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл миграции %s: %v", filePath, err)
	}

	sql := string(sqlBytes)

	// Выполнение SQL запроса
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("не удалось выполнить миграцию %s: %v", filePath, err)
	}

	return nil
}
