package repositories

import (
	"fmt"
	"testing"

	"github.com/Shunkicreate/chi_todo/internal/entities"
)

func TestTodoRepository(t *testing.T) {
    repo := NewTodoRepository()

    // 初期値がない時のGetAll
    todos, err := repo.GetAll()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if len(todos) != 0 {
        t.Fatalf("Expected 0 todos, got %d", len(todos))
    }

    // 初期値がない時のGetByID
    _, err = repo.GetByID(1)
    if err == nil {
        t.Fatalf("Expected error, got none")
    }

    // Create
    todo := entities.Todo{Title: "Test Todo", Completed: false}
    err = repo.Create(todo)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // GetAll
    todos, err = repo.GetAll()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if len(todos) != 1 {
        t.Fatalf("Expected 1 todo, got %d", len(todos))
    }

    // GetByID
    fetchedTodo, err := repo.GetByID(todos[0].ID)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if fetchedTodo.ID != todos[0].ID {
        t.Fatalf("Expected ID %d, got %d", todos[0].ID, fetchedTodo.ID)
    }

    // 存在しないIDでのGetByID
    _, err = repo.GetByID(999)
    if err == nil {
        t.Fatalf("Expected error, got none")
    }

    // Update
    todos[0].Title = "Updated Todo"
    err = repo.Update(todos[0])
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    updatedTodo, err := repo.GetByID(todos[0].ID)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if updatedTodo.Title != "Updated Todo" {
        t.Fatalf("Expected title 'Updated Todo', got '%s'", updatedTodo.Title)
    }

    // 存在しないIDでのUpdate
    err = repo.Update(entities.Todo{ID: 999, Title: "Nonexistent Todo"})
    if err == nil {
        t.Fatalf("Expected error, got none")
    }

    // Delete
    err = repo.Delete(todos[0].ID)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    _, err = repo.GetByID(todos[0].ID)
    if err == nil {
        t.Fatalf("Expected error, got none")
    }

    // 存在しないIDでのDelete
    err = repo.Delete(999)
    if err == nil {
        t.Fatalf("Expected error, got none")
    }

    // Todoが多い時の処理
    for i := 0; i < 100; i++ {
        repo.Create(entities.Todo{Title: "Todo " + fmt.Sprint(i), Completed: false})
    }
    todos, err = repo.GetAll()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if len(todos) != 100 {
        t.Fatalf("Expected 100 todos, got %d", len(todos))
    }
    // 全てのTodoを削除するためのIDを取得
    ids := make([]int, len(todos))
    for i, todo := range todos {
        ids[i] = todo.ID
    }
    // 全てのTodoを削除
    for _, id := range ids {
        err := repo.Delete(id)
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }
    }
    todos, err = repo.GetAll()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if len(todos) != 0 {
        t.Fatalf("Expected 0 todos, got %d", len(todos))
    }
}
