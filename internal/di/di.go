package di

import (
	"github.com/Shunkicreate/chi_todo/internal/interfaces/controllers"
	"github.com/Shunkicreate/chi_todo/internal/interfaces/gateways"
	"github.com/Shunkicreate/chi_todo/internal/repositories"
	"github.com/Shunkicreate/chi_todo/internal/usecases"
	"log"
)

type Dependencies struct {
	TodoUseCase    usecases.TodoUseCase
	TodoRepo       gateways.TodoRepository
	TodoController controllers.TodoController
}

func InitializeDependencies() *Dependencies {
	// MySQLリポジトリの初期化
	dataSourceName := "todo:todopassword@tcp(db:3306)/tododb"
	todoRepo, err := repositories.NewTodoMySQLRepository(dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}

	todoUseCase := usecases.NewTodoUseCase(todoRepo)
	todoController := controllers.TodoController{UseCase: todoUseCase}

	return &Dependencies{
		TodoRepo:       todoRepo,
		TodoUseCase:    todoUseCase,
		TodoController: todoController,
	}
}
