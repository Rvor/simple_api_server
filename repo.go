package main

import (
	"fmt"
	"time"
)

var currentId int

var todos Todos

func init() {
	Create(Todo{Name: "Learning Go"})
	Create(Todo{Name: "Create layout for NOD"})
}

func Create(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	t.Due = time.Now().Add(time.Hour * 3 * 24)
	todos = append(todos, t)
	return t
}

func FindById(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func Destroy(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
