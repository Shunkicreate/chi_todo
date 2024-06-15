package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Shunkicreate/chi_todo/internal/usecases"
)

func TestDeleteTodoHandler(t *testing.T) {
	mockUseCase := new(usecases.MockTodoUseCase)
	todoController := TodoController{UseCase: mockUseCase}

	t.Run("Successful Deletion", func(t *testing.T) {
		todoID := 1
		mockUseCase.On("DeleteTodo", todoID).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/todos/"+strconv.Itoa(todoID), nil)
		recorder := httptest.NewRecorder()

		todoController.DeleteTodo(recorder, req)

		assert.Equal(t, http.StatusNoContent, recorder.Code)
		mockUseCase.AssertCalled(t, "DeleteTodo", todoID)
	})

	t.Run("Deletion with Nonexistent ID", func(t *testing.T) {
		todoID := 999
		mockUseCase.On("DeleteTodo", todoID).Return(fmt.Errorf("todo not found with ID: %d", todoID)).Once()

		req := httptest.NewRequest(http.MethodDelete, "/todos/"+strconv.Itoa(todoID), nil)
		recorder := httptest.NewRecorder()

		todoController.DeleteTodo(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		mockUseCase.AssertCalled(t, "DeleteTodo", todoID)
	})

	t.Run("Invalid ID Format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/todos/invalid", nil)
		recorder := httptest.NewRecorder()

		todoController.DeleteTodo(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
