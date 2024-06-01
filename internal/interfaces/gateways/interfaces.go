package gateways

import "github.com/Shunkicreate/chi_todo/internal/entities"

type TodoRepository interface {
    GetAll() ([]entities.Todo, error)
    GetByID(id int) (entities.Todo, error)
    Create(todo entities.Todo) error
    Update(todo entities.Todo) error
    Delete(id int) error
}
