package todo_test

import (
	"my-first-api/internal/todo"
	"reflect"
	"testing"
)

// TestService_Search tests the Search method of the Service to ensure it properly filters todo items based on a query string.
func TestService_Search(t *testing.T) {
	tests := []struct {
		name           string
		todostoAdd     []string
		query          string
		expectedResult []string
	}{
		// TODO: Add test cases.
		{
			name:           "given a todo of shop and a search of sh, i should get shop back",
			todostoAdd:     []string{"shop"},
			query:          "sh",
			expectedResult: []string{"shop"},
		},
		{
			name:           "still returns shop, even if the case doesnt match",
			todostoAdd:     []string{"Shopping"},
			query:          "sh",
			expectedResult: []string{"Shopping"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := todo.NewService()
			for _, toAdd := range tt.todostoAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Error(err)
				}
			}
			if got := svc.Search(tt.query); !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
