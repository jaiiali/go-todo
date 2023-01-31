package usecases

import (
	"time"

	"github.com/jaiiali/go-todo/helpers"
	"github.com/jaiiali/go-todo/internal/core/domain"
	"github.com/jaiiali/go-todo/internal/core/ports"
	"github.com/jaiiali/go-todo/pkg/logger"
)

type todoUseCase struct {
	repo ports.TodoRepository
	log  *logger.Logger
}

func (t *todoUseCase) FindAll() ([]domain.Todo, error) {
	result, err := t.repo.FindAll()
	if err != nil {
		t.log.Error(err)
		return nil, err
	}

	return result, nil
}

func (t *todoUseCase) FindByID(id string) (*domain.Todo, error) {
	result, err := t.repo.FindByID(id)
	if err != nil {
		t.log.Error(err)
		return nil, err
	}

	return result, nil
}

func (t *todoUseCase) Create(title, description string) (*domain.Todo, error) {
	todo := domain.NewTodo(helpers.NewULID(), title, description, time.Now())

	result, err := t.repo.Create(todo)
	if err != nil {
		t.log.Error(err)
		return nil, err
	}

	return result, nil
}

func NewTodoUseCase(repo ports.TodoRepository, log *logger.Logger) ports.TodoUseCase {
	return &todoUseCase{
		repo: repo,
		log:  log,
	}
}
