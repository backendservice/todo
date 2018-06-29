package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryService_Create(t *testing.T) {
	instance := NewInMemoryService()
	item := Item1()
	err := instance.Create(item)
	assert.NoError(t, err)
}

func TestInMemoryService_Update(t *testing.T) {
	instance := NewInMemoryService()
	item := Item1()
	item.Title = "updated"
	instance.Create(item)
	err := instance.Update(item)
	assert.NoError(t, err)
}

func TestInMemoryService_Update_ErrorNotFound(t *testing.T) {
	instance := NewInMemoryService()
	item := Item2()
	item.Title = "updated"
	err := instance.Update(item)
	assert.Error(t, err)
}
func TestInMemoryService_Delete(t *testing.T) {
	instance := NewInMemoryService()
	item := Item1()
	instance.Create(item)
	err := instance.Delete(item.ID)
	assert.NoError(t, err)
}
