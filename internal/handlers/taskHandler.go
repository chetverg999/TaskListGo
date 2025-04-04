package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"testTask/internal/database"
	"testTask/internal/models"
	"time"
)

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with title, description, and status
// @Param task body models.Task true "Task data"
// @Success 201 {object} models.Task
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /tasks [post]
func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверные данные"})
	}

	if err := task.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка валидации данных", "details": err.Error()})
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	query := `INSERT INTO tasks (title, description, status, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := database.DB.QueryRow(context.Background(), query,
		task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при создании задачи"})
	}

	return c.Status(http.StatusCreated).JSON(task)
}

// GetTasks godoc
// @Summary Get all tasks
// @Description Get a list of all tasks in the database
// @Success 200 {array} models.Task
// @Failure 500 {object} string "Internal server error"
// @Router /tasks [get]
func GetTasks(c *fiber.Ctx) error {
	rows, err := database.DB.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка получения задач"})
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка обработки данных"})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task by ID
// @Param id path int true "Task ID"
// @Param task body models.Task true "Updated task data"
// @Success 200 {object} string "Task updated successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /tasks/{id} [put]
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверные данные"})
	}

	task.UpdatedAt = time.Now()

	query := `UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=$4 WHERE id=$5`
	_, err := database.DB.Exec(context.Background(), query,
		task.Title, task.Description, task.Status, task.UpdatedAt, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка обновления задачи"})
	}

	return c.JSON(fiber.Map{"message": "Задача обновлена"})
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Param id path int true "Task ID"
// @Success 200 {object} string "Task deleted successfully"
// @Failure 500 {object} string "Internal server error"
// @Router /tasks/{id} [delete]
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка удаления задачи"})
	}

	return c.JSON(fiber.Map{"message": "Задача удалена"})
}
