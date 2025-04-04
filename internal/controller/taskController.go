package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"testTask/internal/model"
	"testTask/internal/service"
)

type TaskController struct {
	taskService *service.TaskService
	ctx         context.Context
}

func NewTaskController(ctx context.Context) *TaskController {
	return &TaskController{ctx: ctx, taskService: service.NewTaskService(ctx)}
}

// Create godoc
// @Summary Create a new task
// @Description Create a new task with title, description, and status
// @Param task body model.Task true "Task data"
// @Success 201 {object} model.Task
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /tasks [post]
func (t *TaskController) Create(c *fiber.Ctx) error {
	var task *model.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверные данные"})
	}

	if err := t.taskService.Create(c, task); err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(task)
}

// GetList godoc
// @Summary Get all tasks
// @Description Get a list of all tasks in the database
// @Success 200 {array} model.Task
// @Failure 500 {object} string "Internal server error"
// @Router /tasks [get]
func (t *TaskController) GetList(c *fiber.Ctx) error {
	tasks, err := t.taskService.Get(c)

	if err != nil {
		return err
	}

	return c.JSON(tasks)
}

// Update godoc
// @Summary Update a task
// @Description Update an existing task by ID
// @Param id path int true "Task ID"
// @Param task body model.Task true "Updated task data"
// @Success 200 {object} string "Task updated successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /tasks/{id} [put]
func (t *TaskController) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var task *model.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверные данные"})
	}

	if err := t.taskService.Update(c, task, id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Задача обновлена"})
}

// Delete godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Param id path int true "Task ID"
// @Success 200 {object} string "Task deleted successfully"
// @Failure 500 {object} string "Internal server error"
// @Router /tasks/{id} [delete]
func (t *TaskController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := t.taskService.Delete(c, id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Задача удалена"})
}
