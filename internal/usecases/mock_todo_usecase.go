package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/Shunkicreate/chi_todo/internal/entities"
)

type MockTodoUseCase struct {
	mock.Mock
}

func (m *MockTodoUseCase) GetTodos() ([]entities.Todo, error) {
	args := m.Called()
	return args.Get(0).([]entities.Todo), args.Error(1)
}

func (m *MockTodoUseCase) CreateTodo(todo entities.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoUseCase) UpdateTodo(todo entities.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoUseCase) DeleteTodo(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTodoUseCase) GetTodoByID(id int) (entities.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Todo), args.Error(1)
}
