package repositories

import (
	"github.com/Shunkicreate/chi_todo/internal/entities"
)

type TodoRepositoryImpl struct {
    Todos []entities.Todo
}

func NewTodoRepository() *TodoRepositoryImpl {
    return &TodoRepositoryImpl{
        Todos: []entities.Todo{},
    }
}

func (tr *TodoRepositoryImpl) GetAll() ([]entities.Todo, error) {
    return tr.Todos, nil
}

func (tr *TodoRepositoryImpl) GetByID(id int) (entities.Todo, error) {
    for _, todo := range tr.Todos {
        if todo.ID == id {
            return todo, nil
        }
    }
    return entities.Todo{}, nil
}

func (tr *TodoRepositoryImpl) Create(todo entities.Todo) error {
    tr.Todos = append(tr.Todos, todo)
    return nil
}

func (tr *TodoRepositoryImpl) Update(todo entities.Todo) error {
    for i, t := range tr.Todos {
        if t.ID == todo.ID {
            tr.Todos[i] = todo
            return nil
        }
    }
    return nil
}

func (tr *TodoRepositoryImpl) Delete(id int) error {
    for i, todo := range tr.Todos {
        if todo.ID == id {
            tr.Todos = append(tr.Todos[:i], tr.Todos[i+1:]...)
            return nil
        }
    }
    return nil
}