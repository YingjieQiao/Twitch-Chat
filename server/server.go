package main

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	// initialise logging
	log, _ := zap.NewProduction()
	defer log.Sync()

	// initialise server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1212"
	}

	// initialise router
	r := chi.NewRouter()

	// initialise DB
	db := NewStore()

	// initialise routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		GetKeyValue(w, r, db, log)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		UpdateKeyValue(w, r, db, log)
	})

	log.Info("Starting server", zap.String("port", port))
	log.Error(http.ListenAndServe(":"+port, r).Error())
}
