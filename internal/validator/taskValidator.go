package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"testTask/internal/model"
	"testTask/internal/repository"
)

type TaskValidator struct {
	task       model.Task
	repository *repository.TaskRepository
}

func NewTaskValidator(repository *repository.TaskRepository) *TaskValidator {
	return &TaskValidator{repository: repository}
}

func (t *TaskValidator) Validate(task model.Task) error {
	t.task = task

	return validator.New().Struct(task)
}

func (t *TaskValidator) IssetValidate(id int) fiber.Map {
	count := t.repository.Count(id)

	if count > 0 {
		return nil
	}

	return fiber.Map{"error": "Ошибка обновления задачи"}
}
