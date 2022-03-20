package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// GetKeyValue returns value for associated Key
func GetKeyValue(w http.ResponseWriter, r *http.Request, db *Store, log *zap.Logger) {
	keys, success := r.URL.Query()["key"]
	w.Header().Set("Content-Type", "application/json")

	if !success || len(keys[0]) < 1 {
		log.Error("No key provided")
		err := Error{Message: "No key provided"}
		json.NewEncoder(w).Encode(err)
		return
	}

	value := db.Get(keys[0])
	json.NewEncoder(w).Encode(Result{Key: keys[0], Value: value})
}

// UpdateKeyValue adds key-value pair to database
// if key already exists, overwrite existing value
func UpdateKeyValue(w http.ResponseWriter, r *http.Request, db *Store, log *zap.Logger) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Error("Error parsing request body", zap.String("message", err.Error()))
		err := Error{Message: "Error parsing request body"}
		json.NewEncoder(w).Encode(err)
		return
	}

	db.Put(req.Key, req.Value)
	json.NewEncoder(w).Encode(req)
}
