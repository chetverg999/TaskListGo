package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"sync"
)

type Psql struct {
	db *pgxpool.Pool
}

var (
	instance *Psql
	once     sync.Once
)

func (p *Psql) Db() *pgxpool.Pool {
	return p.db
}

func NewPsql() *Psql {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
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

		instance = &Psql{db: dbpool}
	})

	return instance
}
