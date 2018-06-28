package todo

import "testing"
import "github.com/stretchr/testify/assert"

func TestService_add(t *testing.T) {
	instance := &service{} //instance := new(service)
	orgLen := len(todos)
	instance.Add("something")
	assert.Equal(t, orgLen+1, len(todos))
	assert.Equal(t, "something", todos[0])

	//todos.remove("something")
	//assert.Equal(t, 0, len(todos))
}

func TestService_Add_Today(t *testing.T) {
	instance := &service{} //instance := new(service)
	orgLen := len(todoMap["2018-06-28"])
	instance.Add("breakfast");
	instance.Add("lunch");
	assert.Equal(t, orgLen+2, len(todoMap["2018-06-28"]))
}
