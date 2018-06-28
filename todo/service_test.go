package todo

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestService_Add(t *testing.T){
	instance := &service{}
	orgLen := len(todos)
	instance.Add("whatever")
	assert.Equal(t, orgLen+1, len(todos))
	assert.Equal(t, "whatever", todos[0])
}

func TestService_Add_Today(t *testing.T){
	instance := &service{}
	orgLen := len(todoMap["2018-06-28"])
	instance.Add("breakfast")
	instance.Add("lunch")
	assert.Equal(t, orgLen + 2, len(todoMap["2018-06-28"]))
	delete(todoMap, "2018-06-28")
	for k, v := range todoMap{
		fmt.Println(k, v)
	}
}