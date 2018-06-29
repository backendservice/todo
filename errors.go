package main

import "net/http"

// ServiceError stands for service layer error
type ServiceError struct {
	error
	Code    int
	Message string
}

// NewNotFoundError returns a not found error
func NewNotFoundError(message string) *ServiceError {
	return &ServiceError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
