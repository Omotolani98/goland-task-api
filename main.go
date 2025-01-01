package main

import (
	"github.com/omotolani98/goland-task-api/internal/db"
	"github.com/omotolani98/goland-task-api/internal/todo"
	"github.com/omotolani98/goland-task-api/internal/transport"
	"log"
)

func main() {
	d, err := db.New("doye", "tolani", "todo", "localhost", 5432)
	if err != nil {
		log.Fatal(err)
	}
	svc := todo.NewService(d)
	server := transport.NewServer(svc)

	if err := server.ServeHTTP(); err != nil {
		log.Fatal(svc)
	}
}
