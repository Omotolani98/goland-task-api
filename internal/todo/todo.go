package todo

import (
	"context"
	"errors"
	"fmt"
	"my_first_api/internal/db"
	"strings"
)

type Manager interface {
	InsertItem(ctx context.Context, item db.Item) error
	GetAllItems(ctx context.Context) ([]db.Item, error)
}

type Service struct {
	db Manager
}

func NewService(db Manager) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) Add(todo string) error {
	items, err := svc.GetAll()
	if err != nil {
		return fmt.Errorf("failed to retrieve items: %w", err)
	}

	for _, t := range items {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}

	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "Pending",
	}); err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}

	return nil
}

func (svc *Service) GetAll() ([]db.Item, error) {
	//return svc.todos
	var result []db.Item
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}

	for _, item := range items {
		result = append(result, db.Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}

	return result, nil
}

func (svc *Service) Search(query string) ([]string, error) {
	items, err := svc.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve items: %w", err)
	}

	var result []string
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Task), strings.ToLower(query)) {
			result = append(result, item.Task)
		}
	}

	return result, nil
}
