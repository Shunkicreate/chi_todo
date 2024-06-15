package usecases

import (
	"github.com/Shunkicreate/chi_todo/internal/entities"
)

type TodoUseCase interface {
	GetTodos() ([]entities.Todo, error)
	CreateTodo(todo entities.Todo) error
	UpdateTodo(todo entities.Todo) error
	DeleteTodo(id int) error
	GetTodoByID(id int) (entities.Todo, error)
}
