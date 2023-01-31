package todo

import (
	"fmt"
	"time"

	"github.com/jaiiali/go-todo/internal/core/domain"
)

type todoReq struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"desc" validate:"required"`
}

type todoResp struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"desc"`
	CreatedAt   JSONTime `json:"created_at"`
}

type JSONTime time.Time

func (t JSONTime) String() string {
	return fmt.Sprintf("%q", time.Time(t).Format("2006-01-02 15:04:05"))
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *todoResp) fromDomain(todo *domain.Todo) {
	t.ID = todo.ID
	t.Title = todo.Title
	t.Description = todo.Description
	t.CreatedAt = JSONTime(todo.CreatedAt)
}

// nolint: unused
func (t *todoResp) toDomain() *domain.Todo {
	return &domain.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		CreatedAt:   time.Time(t.CreatedAt),
	}
}

type todoListResp []todoResp

func (t *todoListResp) fromDomain(items []domain.Todo) {
	for _, v := range items {
		todo := todoResp{}
		todo.fromDomain(&v)
		*t = append(*t, todo)
	}
}
