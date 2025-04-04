package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"testTask/internal/model"
	"testTask/internal/repository"
	"testTask/internal/validator"
	"time"
)

type TaskService struct {
	validator  *validator.TaskValidator
	repository *repository.TaskRepository
	ctx        context.Context
}

func NewTaskService(ctx context.Context) *TaskService {
	return &TaskService{ctx: ctx, validator: &validator.TaskValidator{}, repository: repository.NewTaskRepository(ctx)}
}

func (t *TaskService) Create(c *fiber.Ctx, task *model.Task) error {
	if err := t.validateRequest(c, task); err != nil {
		return err
	}

	task.UpdatedAt = time.Now()
	task.CreatedAt = time.Now()

	if err := t.repository.Create(task); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при создании задачи"})
	}

	return nil
}

func (t *TaskService) Update(c *fiber.Ctx, task *model.Task, id int) error {
	if err := t.validateRequest(c, task); err != nil {
		return err
	}

	task.UpdatedAt = time.Now()

	if err := t.repository.Update(task, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка обновления задачи"})
	}

	return nil
}

func (t *TaskService) Delete(c *fiber.Ctx, id int) error {
	if err := t.repository.Delete(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка удаления задачи"})
	}

	return nil
}

func (t *TaskService) Get(c *fiber.Ctx) ([]model.Task, error) {
	rows, err := t.repository.Get()

	defer rows.Close()

	if err != nil {
		return nil, c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка получения задач"})
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task

		if err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка обработки данных"})
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *TaskService) validateRequest(c *fiber.Ctx, task *model.Task) error {
	if err := t.validator.Validate(*task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка валидации данных", "details": err.Error()})
	}

	return nil
}
