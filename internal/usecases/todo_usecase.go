package usecases

import (
	"github.com/Shunkicreate/chi_todo/internal/entities"
	"github.com/Shunkicreate/chi_todo/internal/interfaces/gateways"
)

type TodoUseCase struct {
	TodoRepo gateways.TodoRepository
}

func (tu *TodoUseCase) GetTodos() ([]entities.Todo, error) {
	return tu.TodoRepo.GetAll()
}

func (tu *TodoUseCase) CreateTodo(todo entities.Todo) error {
	return tu.TodoRepo.Create(todo)
}

func (tu *TodoUseCase) UpdateTodo(todo entities.Todo) error {
	return tu.TodoRepo.Update(todo)
}

func (tu *TodoUseCase) DeleteTodo(id int) error {
	return tu.TodoRepo.Delete(id)
}

func (tu *TodoUseCase) GetTodoByID(id int) (entities.Todo, error) {
	return tu.TodoRepo.GetByID(id)
}

