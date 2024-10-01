// entry point of the application
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codersgyan/students-api/internal/config"
)

func main() {

	// load configuration

	cfg := config.MustLoad()

	//database  setup
	// setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Welcome to student API %s", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server")
	}

}
