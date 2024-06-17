package repositories

import (
	"database/sql"
	"github.com/Shunkicreate/chi_todo/internal/entities"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type TodoMySQLRepository struct {
	db *sql.DB
}

func NewTodoMySQLRepository(dataSourceName string) (*TodoMySQLRepository, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &TodoMySQLRepository{db: db}, nil
}

func (r *TodoMySQLRepository) GetAll() ([]entities.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed, link FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []entities.Todo
	for rows.Next() {
		var todo entities.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.Link); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoMySQLRepository) GetByID(id int) (entities.Todo, error) {
	var todo entities.Todo
	err := r.db.QueryRow("SELECT id, title, description, completed, link FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			return todo, fmt.Errorf("todo not found with ID: %d", id)
		}
		return todo, err
	}
	return todo, nil
}

func (r *TodoMySQLRepository) Create(todo entities.Todo) error {
	result, err := r.db.Exec("INSERT INTO todos (title, description, completed, link) VALUES (?, ?, ?, ?)",
		todo.Title, todo.Description, todo.Completed, todo.Link)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.ID = int(id)
	return nil
}

func (r *TodoMySQLRepository) Update(todo entities.Todo) error {
	_, err := r.db.Exec("UPDATE todos SET title = ?, description = ?, completed = ?, link = ? WHERE id = ?",
		todo.Title, todo.Description, todo.Completed, todo.Link, todo.ID)
	if err != nil {
		return fmt.Errorf("todo not found with ID: %d", todo.ID)
	}
	return nil
}

func (r *TodoMySQLRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("todo not found with ID: %d", id)
	}
	return nil
}
