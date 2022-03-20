package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Servers struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Port string `json:"port"`
}

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

	header := r.Header.Get("isBroadcast")
	if header == "1" {
		client := &http.Client{}

		// boardcast to other servers
		jsonFile, err := os.Open("serverConfig.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			log.Error("Error in parsing json config file", zap.String("message", err.Error()))
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var servers Servers
		json.Unmarshal([]byte(byteValue), &servers)

		// build json request body
		data := Request{
			Key:   req.Key,
			Value: req.Value,
		}
		jsonData, _ := json.Marshal(data)

		if err != nil {
			log.Error("Error in creating request body", zap.String("message", err.Error()))
		}

		for i := 0; i < len(servers.Servers); i++ {
			targetPort := servers.Servers[i].Port
			currPort := os.Getenv("PORT")
			if currPort == targetPort {
				continue
			}

			log.Info("Updating server", zap.String("port", targetPort))
			endpoint := fmt.Sprintf("http://host.docker.internal:%s/", targetPort)
			broadcastRequest, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
			//_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
			broadcastRequest.Header.Add("isBroadcast", "0")
			broadcastRequest.Header.Add("Content-Type", "application/json")
			if err != nil {
				log.Error("Request Failed", zap.String("message", err.Error()))
				return
			}
			resp, _ := client.Do(broadcastRequest)
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			log.Info("Update server status", zap.String("response body", bodyString))
		}
	}

	json.NewEncoder(w).Encode(req)
}
