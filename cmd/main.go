package main

import (
	"fmt"
	"log"
	// "net/http"
	// "todo-app/internal/frameworks/webserver"
	// "todo-app/internal/interfaces/controllers"
	// "todo-app/internal/usecases"
	// "todo-app/internal/repositories"
)

func main() {
    // todoRepo := repositories.NewTodoRepository() // 仮実装
    // todoUseCase := usecases.TodoUseCase{TodoRepo: todoRepo}
    // todoController := controllers.TodoController{UseCase: todoUseCase}

    // router := webserver.NewRouter(todoController)

    log.Println("Server started at :8080")
    fmt.Println("Hello, World!")
    // if err := http.ListenAndServe(":8080", router); err != nil {
    //     log.Fatalf("could not listen on port 8080 %v", err)
    // }
}
