package repositories

import (
	"github.com/Shunkicreate/chi_todo/internal/entities"
	"sync"
    "fmt"
)

type TodoRepositoryImpl struct {
	todos  []entities.Todo
	lastID int
	mu     sync.Mutex
}

// NewTodoRepository は新しい TodoRepositoryImpl インスタンスを作成して返します。
// この関数は、空の Todo スライスを持つ TodoRepositoryImpl のポインタを返します。
func NewTodoRepository() *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		todos: []entities.Todo{},
        lastID: 0,
	}
}

func (tr *TodoRepositoryImpl) GetAll() ([]entities.Todo, error) {
    tr.mu.Lock()
    defer tr.mu.Unlock()
    return tr.todos, nil
}

func (tr *TodoRepositoryImpl) GetByID(id int) (entities.Todo, error) {
    tr.mu.Lock()
    defer tr.mu.Unlock()
    for _, todo := range tr.todos {
        if todo.ID == id {
            return todo, nil
        }
    }
    return entities.Todo{}, fmt.Errorf("todo not found with ID: %d", id)
}

func (tr *TodoRepositoryImpl) Create(todo entities.Todo) error {
    tr.mu.Lock()
    defer tr.mu.Unlock()
    tr.lastID++
    todo.ID = tr.lastID
    tr.todos = append(tr.todos, todo)
    return nil
}

func (tr *TodoRepositoryImpl) Update(todo entities.Todo) error {
    tr.mu.Lock()
    defer tr.mu.Unlock()
    for i, t := range tr.todos {
        if t.ID == todo.ID {
            tr.todos[i] = todo
            return nil
        }
    }
    return fmt.Errorf("todo not found with ID: %d", todo.ID)
}

// Delete は指定されたIDのTODOを削除します。
// もし指定されたIDのTODOが存在しない場合は、エラーを返します。
//
// パラメータ:
//   - id: 削除するTODOのID
//
// 戻り値:
//   - error: エラーが発生した場合はその情報を返します。削除が成功した場合はnilを返します。
func (tr *TodoRepositoryImpl) Delete(id int) error {
    tr.mu.Lock()
    defer tr.mu.Unlock()
    for i, todo := range tr.todos {
        if todo.ID == id {
            tr.todos = append(tr.todos[:i], tr.todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("todo not found with ID: %d", id)
}