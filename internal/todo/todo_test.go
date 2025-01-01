package todo

import (
	"context"
	"github.com/omotolani98/goland-task-api/internal/db"
	"reflect"
	"testing"
)

type MockDB struct {
	items []db.Item
}

func (m *MockDB) InsertItem(_ context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(_ context.Context) ([]db.Item, error) {
	return m.items, nil
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		name               string
		toDosToAdd         []string
		query              string
		expectedTestResult []string
	}{
		{
			name:               "given a todo of grocery and a search of sh, I should get grocery back",
			toDosToAdd:         []string{"grocery"},
			query:              "groc",
			expectedTestResult: []string{"grocery"},
		},
		{
			name:               "still returns grocery, even when case does not match",
			toDosToAdd:         []string{"grocery"},
			query:              "Groc",
			expectedTestResult: []string{"grocery"},
		},
		{
			name:               "spaces",
			toDosToAdd:         []string{"go grocery shopping"},
			query:              "Groc",
			expectedTestResult: []string{"go grocery shopping"},
		}, {
			name:               "space at the start of the word",
			toDosToAdd:         []string{" go grocery shopping"},
			query:              "Groc",
			expectedTestResult: []string{" go grocery shopping"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := NewService(m)
			for _, toAdd := range tt.toDosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Errorf("%v", err)
				}
			}
			got, err := svc.Search(tt.query)
			if err != nil {
				t.Errorf("%v", err)
			}
			if !reflect.DeepEqual(got, tt.expectedTestResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedTestResult)
			}
		})
	}
}
