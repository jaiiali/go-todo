package ports

import (
	"github.com/jaiiali/go-todo/internal/core/domain"
)

type TodoUseCase interface {
	FindAll() ([]domain.Todo, error)
	FindByID(id string) (*domain.Todo, error)
	Create(title, description string) (*domain.Todo, error)
}
