package webserver

import (
	"github.com/Shunkicreate/chi_todo/internal/interfaces/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(todoController controllers.TodoController) http.Handler {
    r := chi.NewRouter()

    // ミドルウェアの設定
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // 静的ファイルの提供
    r.Handle("/", http.FileServer(http.Dir("./internal/frameworks/webserver/static")))

    // APIルートの設定
    r.Get("/todos", todoController.GetTodos)
    // r.Get("/todos/{id}", todoController.GetTodoByID)
    r.Post("/todos", todoController.CreateTodo)
    r.Put("/todos/{id}", todoController.UpdateTodo)
    r.Delete("/todos/{id}", todoController.DeleteTodo)

    return r
}