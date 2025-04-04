package main

import (
	_ "awesomeProject2/docs"
	"awesomeProject2/internal/database"
	"awesomeProject2/internal/handlers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title           TODO List API
// @version         1.0
// @description     API для работы с задачами
// @BasePath        /api
func main() {
	database.ConnectDB()

	app := fiber.New()

	// Эндпоинт для документации Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	// API маршруты
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks", handlers.GetTasks)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	fmt.Println("🚀 Сервер запущен на http://localhost:3000")
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
