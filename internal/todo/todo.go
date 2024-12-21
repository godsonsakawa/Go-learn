package todo

import (
	"context"
	"errors"
	"fmt"
	"my-first-api/internal/db"
	"strings"
)

// Item represents a task with its associated status in the todo list.
type Item struct {
	Task   string
	Status string
}

// Service provides operations for managing a collection of todo items.
type Service struct {
	db *db.DB
}

// generate a constructor

// NewService initializes a new Service instance with an empty TODO list and returns a pointer to it.
func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

// Add a TODO Method

// Add adds a new todo item to the service if it is unique and returns an error if the item already exists.
func (svc *Service) Add(todo string) error {
	items, err := svc.GetAll()
	if err != nil {
		return fmt.Errorf("could not get all items: %w", err)
	}

	for _, t := range items {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	}); err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}
	return nil
}

// Search filters and returns a list of todo items containing the specified query string.
func (svc *Service) Search(query string) ([]string, error) {
	items, err := svc.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not get all items: %w", err)
	}

	var results []string
	for _, todo := range items {
		if strings.Contains(strings.ToLower(todo.Task), strings.ToLower(query)) {
			results = append(results, todo.Task)
		}
	}
	return results, nil
}

// GetAll retrieves all items managed by the Service instance. It returns a slice of Item objects.
func (svc *Service) GetAll() ([]Item, error) {
	var results []Item
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not read from db: %w", err)
	}
	for _, item := range items {
		results = append(results, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}
	return results, nil
}
