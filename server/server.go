package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	// initialise server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1212"
	}

	// initialise router
	r := chi.NewRouter()

	// initialise routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Printf("Running on port [%s]", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
