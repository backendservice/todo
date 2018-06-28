package todo

// TodoService defines what to do with this service
type Service interface {
	Add(item string)
}