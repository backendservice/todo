package main

import (
	"time"
)

// Service defines what todo list service can do
type Service interface {
	Create(Item) error
	List() []Item
	Update(Item) error
	Delete(string) error
}

type inMemoryService struct {
	todos       map[string]Item
	sortedTodos []Item
}

// NewInMemoryService provides a instance of in-memory implementation
func NewInMemoryService() Service {
	return &inMemoryService{
		todos:       map[string]Item{},
		sortedTodos: []Item{},
	}
}

func (s *inMemoryService) Create(item Item) error {
	item.CreatedAt = time.Now()
	s.todos[item.ID] = item
	// s.sortedTodos = append(s.sortedTodos, item)
	return nil
}

func (s *inMemoryService) List() []Item {

	var items []Item
	for k := range s.todos {
		items = append(items, s.todos[k])
	}
	return items
}

func (s *inMemoryService) Update(item Item) error {
	if _, ok := s.todos[item.ID]; ok {
		s.todos[item.ID] = item
	}

	return nil
}

func (s *inMemoryService) Delete(id string) error {
	if _, ok := s.todos[id]; ok {
		delete(s.todos, id)
	}

	return nil
}
