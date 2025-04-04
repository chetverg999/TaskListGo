package validator

import (
	"github.com/go-playground/validator/v10"
	"testTask/internal/model"
)

type TaskValidator struct {
	task model.Task
}

func (t *TaskValidator) Validate(task model.Task) error {
	t.task = task

	return validator.New().Struct(task)
}
