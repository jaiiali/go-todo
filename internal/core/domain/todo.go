package domain

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
}

func NewTodo(id, title, description string, createdAt time.Time) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func (t *Todo) String() string {
	return fmt.Sprintf("%s: %s\n%s\n%s", t.ID, t.Title, t.Description, t.CreatedAt)
}
