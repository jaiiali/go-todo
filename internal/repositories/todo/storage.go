package todo

import (
	"time"

	"github.com/jaiiali/go-todo/internal/core/domain"
)

type todoMemoryStorage struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
}

func (s *todoMemoryStorage) fromDomain(todo *domain.Todo) {
	s.ID = todo.ID
	s.Title = todo.Title
	s.Description = todo.Description
	s.CreatedAt = todo.CreatedAt
}

func (s *todoMemoryStorage) toDomain() *domain.Todo {
	return &domain.Todo{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
	}
}
