package main

import (
	"log"
	"net/http"

	"github.com/Shunkicreate/chi_todo/internal/di"
	"github.com/Shunkicreate/chi_todo/internal/frameworks/webserver"
)

func main() {
	deps := di.InitializeDependencies()
	router := webserver.NewRouter(deps.TodoController)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
