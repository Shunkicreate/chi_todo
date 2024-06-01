package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shunkicreate/chi_todo/internal/usecases"
)

type TodoController struct {
	UseCase usecases.TodoUseCase
}

func (tc *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.UseCase.GetTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
