package main

import (
	"time"
)

// Service defines what todo list service can do
type Service interface {
	Create(Item) error
}

type inMemoryService struct {
	todos map[string]Item
}

// NewInMemoryService provides a instance of in-memory implementation
func NewInMemoryService() Service {
	return &inMemoryService{
		todos: map[string]Item{},
	}
}

func (s *inMemoryService) Create(item Item) error {
	item.CreatedAt = time.Now()
	s.todos[item.ID] = item
	return nil
}
