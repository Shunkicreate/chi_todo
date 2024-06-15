package usecases

import (
	"github.com/Shunkicreate/chi_todo/internal/entities"
	"github.com/Shunkicreate/chi_todo/internal/interfaces/gateways"
)

type TodoUseCaseImpl struct {
	TodoRepo gateways.TodoRepository
}

func NewTodoUseCase(todoRepo gateways.TodoRepository) TodoUseCase {
	return &TodoUseCaseImpl{TodoRepo: todoRepo}
}

func (tu *TodoUseCaseImpl) GetTodos() ([]entities.Todo, error) {
	return tu.TodoRepo.GetAll()
}

func (tu *TodoUseCaseImpl) CreateTodo(todo entities.Todo) error {
	return tu.TodoRepo.Create(todo)
}

func (tu *TodoUseCaseImpl) UpdateTodo(todo entities.Todo) error {
	return tu.TodoRepo.Update(todo)
}

func (tu *TodoUseCaseImpl) DeleteTodo(id int) error {
	return tu.TodoRepo.Delete(id)
}

func (tu *TodoUseCaseImpl) GetTodoByID(id int) (entities.Todo, error) {
	return tu.TodoRepo.GetByID(id)
}
