package todo

import (
	"errors"
	"strings"
)

// Item represents a task with its associated status in the todo list.
type Item struct {
	Task   string
	Status string
}

// Service provides operations for managing a collection of todo items.
type Service struct {
	todos []Item
}

// generate a constructor

// NewService initializes a new Service instance with an empty TODO list and returns a pointer to it.
func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

// Add a TODO Method

// Add adds a new todo item to the service if it is unique and returns an error if the item already exists.
func (svc *Service) Add(todo string) error {
	for _, t := range svc.todos {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	svc.todos = append(svc.todos, Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	})
	return nil
}

// Search filters and returns a list of todo items containing the specified query string.
func (svc *Service) Search(query string) []string {
	var results []string
	for _, todo := range svc.todos {
		if strings.Contains(strings.ToLower(todo.Task), strings.ToLower(query)) {
			results = append(results, todo.Task)
		}
	}
	return results
}

// GetAll retrieves all items managed by the Service instance. It returns a slice of Item objects.
func (svc *Service) GetAll() []Item {
	return svc.todos
}
