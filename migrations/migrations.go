package migrations

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	"os"
	"path/filepath"
	"testTask/internal/database"
)

type Migration struct {
	psql *database.Psql
}

func NewMigration(psql *database.Psql) *Migration {
	return &Migration{psql: psql}
}

func (m *Migration) MigrateDatabase(migrationDir string) {
	files, err := os.ReadDir(migrationDir)

	if err != nil {
		log.Fatalf("Ошибка при чтении директории миграций: %v\n", err)
	}

	for _, file := range files {
		if file.IsDir() || !m.isSQLFile(file.Name()) {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", migrationDir, file.Name())

		if err = m.applyMigration(m.psql.Db(), filePath); err != nil {
			log.Fatalf("Ошибка при применении миграции %s: %v\n", filePath, err)
		}
	}
}

func (m *Migration) isSQLFile(fileName string) bool {

	return filepath.Ext(fileName) == ".sql"
}

func (m *Migration) applyMigration(conn *pgxpool.Pool, filePath string) error {
	sqlBytes, err := os.ReadFile(filePath)

	if err != nil {
		return fmt.Errorf("не удалось прочитать файл миграции %s: %v", filePath, err)
	}

	if _, err = conn.Exec(context.Background(), string(sqlBytes)); err != nil {
		return fmt.Errorf("не удалось выполнить миграцию %s: %v", filePath, err)
	}

	return nil
}
