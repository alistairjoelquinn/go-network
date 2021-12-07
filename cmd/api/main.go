package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {
	port int
}

func main() {
	app := &application{
		port: 3001,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      app.routes(),
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	log.Println("Server running on port", app.port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
