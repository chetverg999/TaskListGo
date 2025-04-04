package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"os"
	"testTask/internal/controller"
	"testTask/internal/database"
	"testTask/internal/helper"
	"testTask/migrations"
)

// @title           API TaskList
// @version         1.0
// @description     API для работы с задачами
// @BasePath        /api
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	psql := database.NewPsql()

	defer psql.Db().Close()

	migrations.NewMigration(psql).MigrateDatabase(os.Getenv("MIGRATION_DIR"))

	app := fiber.New()

	taskController := controller.NewTaskController(context.WithValue(context.Background(), helper.PsqlKey, psql))

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})
	app.Post("/tasks", taskController.Create)
	app.Get("/tasks", taskController.GetList)
	app.Put("/tasks/:id", taskController.Update)
	app.Delete("/tasks/:id", taskController.Delete)

	if err = app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err)
	}
}
