package repositories

import "github.com/Shunkicreate/chi_todo/internal/entities"

func (r *TodoMySQLRepository) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		link TEXT,
		PRIMARY KEY (id)
	);`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	// 初期データの挿入
	initialTodos := []entities.Todo{
		{Title: "Todo 1", Description: "Description 1", Completed: false, Link: "http://example.com/1"},
		{Title: "Todo 2", Description: "Description 2", Completed: false, Link: "http://example.com/2"},
		{Title: "Todo 3", Description: "Description 3", Completed: false, Link: "http://example.com/3"},
		{Title: "Todo 4", Description: "Description 4", Completed: false, Link: "http://example.com/4"},
		{Title: "Todo 5", Description: "Description 5", Completed: false, Link: "http://example.com/5"},
		{Title: "Todo 6", Description: "Description 6", Completed: false, Link: "http://example.com/6"},
		{Title: "Todo 7", Description: "Description 7", Completed: false, Link: "http://example.com/7"},
		{Title: "Todo 8", Description: "Description 8", Completed: false, Link: "http://example.com/8"},
		{Title: "Todo 9", Description: "Description 9", Completed: false, Link: "http://example.com/9"},
		{Title: "Todo 10", Description: "Description 10", Completed: false, Link: "http://example.com/10"},
	}

	for _, todo := range initialTodos {
		if err := r.Create(todo); err != nil {
			return err
		}
	}

	return nil
}