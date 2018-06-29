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
	todos map[string]Item
	ids []string
}

// NewInMemoryService provides a instance of in-memory implementation
func NewInMemoryService() Service {
	return &inMemoryService{
		todos: map[string]Item{},
		ids: []string{},
	}
}

func (s *inMemoryService) Create(item Item) error {
	item.CreatedAt = time.Now()
	s.todos[item.ID] = item
	s.ids = append(s.ids, item.ID)
	return nil
}

func (s *inMemoryService) List() []Item {
	var list []Item
	for _, id := range s.ids {
		if item, ok := s.todos[id]; ok {
			list = append(list, item)
		}
	}
	return list
}

func (s *inMemoryService) Update(item Item) error {
	if _, ok := s.todos[item.ID]; ok {
		s.todos[item.ID] = item
	}else{
		return NewNotFoundError("Not exist item.")
	}
	return nil
}

func (s *inMemoryService) Delete(itemId string) error {
	delete(s.todos, itemId)
	return nil
}