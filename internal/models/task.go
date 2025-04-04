package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status" validate:"required,oneof=new in_progress done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
