package main

import "time"

// Item is a line item of todo list
type Item struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"-"`
}
