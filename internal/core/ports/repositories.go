package ports

import (
	"github.com/jaiiali/go-todo/internal/core/domain"
)

type TodoRepository interface {
	FindAll() ([]domain.Todo, error)
	FindByID(id string) (*domain.Todo, error)
	Create(todo *domain.Todo) (*domain.Todo, error)
}
