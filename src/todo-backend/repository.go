package main

import (
	"fmt"
	"sync"
)

// TodoService defines an interface for the data methods to support different storage types
type TodoService interface {
	GetAll() ([]Todo, error)
	Get(id int) (*Todo, error)
	Save(todo *Todo) error
	DeleteAll() error
	Delete(id int) error
}

// MockTodoService uses a concurrent array for basic testing
type MockTodoService struct {
	m      sync.Mutex
	nextID int
	Todos  []*Todo
}

// NewMockTodoService creates a mocked todo service
func NewMockTodoService() *MockTodoService {
	t := new(MockTodoService)
	t.m.Lock()
	t.Todos = make([]*Todo, 0)
	t.nextID = 1 // Start at 1 so we can distinguish from unspecified (0)
	t.m.Unlock()
	return t
}

// GetAll returns all todos
func (t *MockTodoService) GetAll() ([]*Todo, error) {
	return t.Todos, nil
}

// Get returns a todo identified by id
func (t *MockTodoService) Get(id int) (*Todo, error) {
	for _, value := range t.Todos {
		if value.ID == id {
			return value, nil
		}
	}
	return nil, nil
}

// Save adds the todo to the repository
func (t *MockTodoService) Save(todo *Todo) error {
	if todo.ID == 0 { // Insert
		t.m.Lock()
		todo.ID = t.nextID
		t.nextID++
		t.m.Unlock()

		t.m.Lock()
		t.Todos = append(t.Todos, todo)
		t.m.Unlock()
		return nil
	}

	// Update existing
	for i, value := range t.Todos {
		if value.ID == todo.ID {
			t.Todos[i] = todo
			return nil
		}
	}

	return fmt.Errorf("Not Found")
}

// DeleteAll deletes all todo entries
func (t *MockTodoService) DeleteAll() error {
	t.m.Lock()
	t.Todos = make([]*Todo, 0)
	t.m.Unlock()
	return nil
}

// Delete deletes the todo identified by id
func (t *MockTodoService) Delete(id int) error {
	for i, value := range t.Todos {
		if value.ID == id {
			t.m.Lock()
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			t.m.Unlock()
			return nil
		}
	}
	return nil
}
