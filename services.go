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
	indexer map[string]int
	todos   []Item
}

// NewInMemoryService provides a instance of in-memory implementation
func NewInMemoryService() Service {
	return &inMemoryService{
		todos:   []Item{},
		indexer: map[string]int{},
	}
}

func (s *inMemoryService) Create(item Item) error {
	item.CreatedAt = time.Now()
	s.todos = append(s.todos, item)
	s.indexer[item.ID] = len(s.todos) - 1
	return nil
}

func (s *inMemoryService) List() ([]Item, error) {
	return s.todos, nil
}

func (s *inMemoryService) Update(item Item) error {
	index := s.indexer[item.ID]
	s.todos[index] = item
	return nil
}

func (s *inMemoryService) Delete(id string) error {
	index := s.indexer[id]
	s.todos = append(s.todos[:index], s.todos[index+1:]...)
	delete(s.indexer, id)
	return nil
}
