package main

import (
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
	todos   map[string]Item
	todoSeq []string
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
	s.todoSeq = append(s.todoSeq, item.ID)
	return nil
}

func (s *inMemoryService) List() ([]Item, error) {
	var result []Item
	for _, k := range s.todoSeq {
		result = append(result, s.todos[k])
	}
	return result, nil
}

func (s *inMemoryService) Update(item Item) error {
	if _, ok := s.todos[item.ID]; ok {
		s.todos[item.ID] = item
	} else {
		return NewNotFoundError("Can't be updated")
	}
	return nil
}

func (s *inMemoryService) Delete(id string) error {
	delete(s.todos, id)
	return nil
}
