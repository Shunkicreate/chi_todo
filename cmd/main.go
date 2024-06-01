package main

import (
	"log"
	"net/http"

	"github.com/Shunkicreate/chi_todo/internal/frameworks/webserver"
	"github.com/Shunkicreate/chi_todo/internal/interfaces/controllers"
	"github.com/Shunkicreate/chi_todo/internal/repositories"
	"github.com/Shunkicreate/chi_todo/internal/usecases"
)

func main() {
	todoRepo := repositories.NewTodoRepository() // 仮実装
	todoUseCase := usecases.TodoUseCase{TodoRepo: todoRepo}
	todoController := controllers.TodoController{UseCase: todoUseCase}

	router := webserver.NewRouter(todoController)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
