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
