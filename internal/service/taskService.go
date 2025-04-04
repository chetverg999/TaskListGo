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
	repository := repository.NewTaskRepository(ctx)

	return &TaskService{ctx: ctx, validator: validator.NewTaskValidator(repository), repository: repository}
}

func (t *TaskService) Create(c *fiber.Ctx, task *model.Task) error {
	if err := t.validateRequest(task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	task.UpdatedAt = time.Now()
	task.CreatedAt = time.Now()

	if err := t.repository.Create(task); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при создании задачи"})
	}

	return c.Status(http.StatusCreated).JSON(task)
}

func (t *TaskService) Update(c *fiber.Ctx, task *model.Task, id int) error {
	if err := t.validateRequest(task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := t.validator.IssetValidate(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	task.ID = id
	task.UpdatedAt = time.Now()

	if err := t.repository.Update(task, task.ID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка обновления задачи"})
	}

	return c.Status(http.StatusOK).JSON(task)
}

func (t *TaskService) Delete(c *fiber.Ctx, id int) error {
	if err := t.validator.IssetValidate(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := t.repository.Delete(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка удаления задачи"})
	}

	return c.Status(http.StatusNoContent).JSON(id)
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

func (t *TaskService) validateRequest(task *model.Task) fiber.Map {
	if err := t.validator.Validate(*task); err != nil {
		return fiber.Map{"error": "Ошибка валидации данных", "details": err.Error()}
	}

	return nil
}
