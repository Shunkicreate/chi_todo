package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Shunkicreate/chi_todo/internal/entities"
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

func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo entities.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := tc.UseCase.CreateTodo(todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (tc *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo entities.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := tc.UseCase.UpdateTodo(todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteTodo はTodoを削除する
// クエリパラメータのidを取得し、それを元にTodoを削除する
func (tc *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr, err := extractIDFromURL(r.URL)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := tc.UseCase.DeleteTodo(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func extractIDFromURL(u *url.URL) (string, error) {
	segments := strings.Split(u.Path, "/")
	if len(segments) < 3 {
		return "", errors.New("invalid URL path")
	}
	return segments[2], nil
}
