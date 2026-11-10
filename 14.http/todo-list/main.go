package main

import (
	"fmt"

	"github.com/seriiserii825/go-course/14.http/todo-list/http"
	"github.com/seriiserii825/go-course/14.http/todo-list/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHttpHandlers(todoList)
	httpServer := http.NewHttpServer(httpHandlers)
	if err := httpServer.Start(); err != nil {
		fmt.Printf("Failed to start HTTP server: %v\n", err)
	}
}
