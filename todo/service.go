package todo

import (
	"time"
)

var todos []string
var todoMap = map[string][]string{}

type service struct {
}

func (s *service) Add(item string) {
	todos = append(todos, item)
	date := time.Now().Format("2006-01-02")
	if list, ok := todoMap[date]; ok {
		todoMap[date] = append(list, item)
	} else {
		todoMap[date] = []string{item}
	}
}
