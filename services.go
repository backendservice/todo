package main

import (
	"fmt"
	"time"
)

// Service defines what todo list service can do
type Service interface {
	Create(Item) error
	List() ([]Item, error)
	Update(Item) error
	Delete(string) error
}

type inMemoryService struct {
	todos       map[string]Item
	sortedTodos []string
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
	s.sortedTodos = append(s.sortedTodos, item.ID)
	return nil
}

func (s *inMemoryService) List() ([]Item, error) {

	var items []Item
	for _, k := range s.sortedTodos {
		items = append(items, s.todos[k])
	}
	return items, nil
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

	for k, v := range s.todos {
		fmt.Println(k, v)
	}
	return nil
}
