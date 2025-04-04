package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"testTask/internal/database"
	"testTask/internal/helper"
	"testTask/internal/model"
)

type TaskRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewTaskRepository(ctx context.Context) *TaskRepository {
	return &TaskRepository{ctx: ctx, db: helper.FindContext(ctx, helper.PsqlKey).(*database.Psql).Db()}
}

func (repository *TaskRepository) Create(task *model.Task) error {
	query := `INSERT INTO tasks (title, description, status, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`

	return repository.db.
		QueryRow(
			context.Background(),
			query,
			task.Title,
			task.Description,
			task.Status,
			task.CreatedAt,
			task.UpdatedAt,
		).
		Scan(&task.ID)
}

func (repository *TaskRepository) Update(task *model.Task, id int) error {
	query := `UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=$4 WHERE id=$5`
	_, err := repository.db.Exec(
		context.Background(),
		query,
		task.Title,
		task.Description,
		task.Status,
		task.UpdatedAt,
		id,
	)

	return err
}

func (repository *TaskRepository) Delete(id int) error {
	_, err := repository.db.Exec(context.Background(), `DELETE FROM tasks WHERE id=$1`, id)

	return err
}

func (repository *TaskRepository) Get() (pgx.Rows, error) {
	return repository.db.Query(context.Background(), `SELECT * FROM tasks`)
}

func (repository *TaskRepository) Count(id int) int {
	var count int
	_ = repository.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM tasks WHERE id=$1`, id).Scan(&count)

	return count
}
