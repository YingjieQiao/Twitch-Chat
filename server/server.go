package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		endpoint := fmt.Sprintf("http://localhost:%s/listen", targetPort)
		fmt.Println(endpoint)
		resp, err := http.Get(endpoint)
		if err != nil {
			log.Printf("Request Failed: %s", err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Reading body failed: %s", err)
			return
		}
		// Log the request body
		bodyString := string(body)
		log.Print(bodyString)
	//case "POST":
	//	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	//	if err := r.ParseForm(); err != nil {
	//		fmt.Fprintf(w, "ParseForm() err: %v", err)
	//		return
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//
	//	str := fmt.Sprintf("Another server on port [%s]", port)
	//	fmt.Printf(str)
	//	data := map[string]string{
	//		"status": "OK",
	//		"info":   str,
	//	}
	//	jsonData, _ := json.Marshal(data)
	//	w.Write(jsonData)

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
		fmt.Printf(str)
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
