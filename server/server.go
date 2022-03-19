package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	switch r.Method {

	case "GET":
		targetPort := r.URL.Query().Get("port")
		endpoint := fmt.Sprintf("http://host.docker.internal:%s/listen", targetPort)
		fmt.Println(endpoint)
		resp, err := http.Get(endpoint)
		if err != nil {
			log.Printf("Request Failed: %s", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println(resp)

	default:
		fmt.Fprintf(w, "Method not supported")
	}
}

func listen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	switch r.Method {

	case "GET":
		str := fmt.Sprintf("Current server on port [%s]", port)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data := map[string]string{
			"status": "OK",
			"info":   str,
		}
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)

	default:
		fmt.Fprintf(w, "Method not supported")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/listen", listen)

	fmt.Printf("Running on port [%s]", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
