package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// 	_ = r.Close
// }

type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handleJSONPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Print received data
	fmt.Printf("Received JSON: %+v\n", data)

	// Send a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", handleJSONPost) // Register the handler
	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	http.ListenAndServe(port, nil) // Start the server
}
