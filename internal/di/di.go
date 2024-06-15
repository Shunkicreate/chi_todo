package di

import (
	"github.com/Shunkicreate/chi_todo/internal/interfaces/controllers"
	"github.com/Shunkicreate/chi_todo/internal/interfaces/gateways"
	"github.com/Shunkicreate/chi_todo/internal/repositories"
	"github.com/Shunkicreate/chi_todo/internal/usecases"
)

type Dependencies struct {
	TodoUseCase    usecases.TodoUseCase
	TodoRepo       gateways.TodoRepository
	TodoController controllers.TodoController
}

func InitializeDependencies() *Dependencies {
	todoRepo := repositories.NewTodoRepository()
	todoUseCase := usecases.NewTodoUseCase(todoRepo)
	todoController := controllers.TodoController{UseCase: todoUseCase}

	return &Dependencies{
		TodoRepo:       todoRepo,
		TodoUseCase:    todoUseCase,
		TodoController: todoController,
	}
}
