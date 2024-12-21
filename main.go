package main

import (
	"log"
	"my-first-api/internal/db"
	"my-first-api/internal/todo"
	"my-first-api/internal/transport"
)

// main is the entry point of the application, initializing and starting the TODO API server.
func main() {
	// A simple API to allow you to create and retrieve a TODO list

	d, err := db.New("postgres", "example", "postgres", "localhost", 5432)
	if err != nil {
		log.Fatal(err)
	}

	svc := todo.NewService(d)          // initializes a new **TODO service** by calling the `NewService` function from the `todo` package, and assigns the resulting pointer to a variable named `svc`.
	server := transport.NewServer(svc) // pass it to the server

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
