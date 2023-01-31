package todo

import (
	"errors"

	"github.com/jaiiali/go-todo/internal/core/domain"
	"github.com/jaiiali/go-todo/pkg/logger"
)

type MemoryRepo struct {
	data map[string]todoMemoryStorage
	log  *logger.Logger
}

func NewMemoryRepo(log *logger.Logger) *MemoryRepo {
	return &MemoryRepo{
		data: make(map[string]todoMemoryStorage),
		log:  log,
	}
}

func (r *MemoryRepo) FindAll() ([]domain.Todo, error) {
	var result = []domain.Todo{}
	for _, v := range r.data {
		result = append(result, *v.toDomain())
	}

	return result, nil
}

func (r *MemoryRepo) FindByID(id string) (*domain.Todo, error) {
	if result, ok := r.data[id]; ok {
		return result.toDomain(), nil
	}

	err := errors.New("todo not found")
	r.log.Error(err)

	return nil, err
}

func (r *MemoryRepo) Create(todo *domain.Todo) (*domain.Todo, error) {
	var storage = &todoMemoryStorage{}
	storage.fromDomain(todo)

	r.data[todo.ID] = *storage

	return todo, nil
}
