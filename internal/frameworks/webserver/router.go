package webserver

import (
	"github.com/Shunkicreate/chi_todo/internal/interfaces/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(todoController controllers.TodoController) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to chi_todo"))
	})
	router.Get("/todos", todoController.GetTodos)

	return router
}
