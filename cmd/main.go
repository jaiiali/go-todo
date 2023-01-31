package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jaiiali/go-todo/helpers"
	"github.com/jaiiali/go-todo/internal/core/usecases"
	handlerTodo "github.com/jaiiali/go-todo/internal/handlers/todo"
	repoTodo "github.com/jaiiali/go-todo/internal/repositories/todo"
	"github.com/jaiiali/go-todo/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	defer log.Sync() //nolint: errcheck

	// Repository
	todoRepo := repoTodo.NewMemoryRepo(log)

	// UseCase
	todoUseCase := usecases.NewTodoUseCase(todoRepo, log)

	app := fiber.New()
	app.Use(recover.New())
	api := app.Group("/api")

	// Handler
	handlerTodo.NewHandler(todoUseCase, api)

	log.Info("Listening...")
	log.Panic(app.Listen(helpers.BuildAddr()))
}
