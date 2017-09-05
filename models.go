package main

import (
	"time"
)

type TodoItem struct {
	Id      string `gorethink:"id,omitempty"`
	Title   string
	Project string
	Done    bool
	Created time.Time
}

func (t *TodoItem) Completed() bool {
	return t.Done
}

func NewTodoItem(title string, project string) *TodoItem {
	return &TodoItem{
		Title:   title,
		Done:    false,
		Project: project,
	}
}
