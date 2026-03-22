package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	
	localVar := "Hello from local variable!"
	envVar := os.Getenv("MY_VAR")
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Local variable: %s\n", localVar)
		fmt.Fprintf(w, "Environment variable MY_VAR: %s\n", envVar)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	
	log.Printf("Server starting on port %s", port)
	log.Printf("Local variable value: %s", localVar)
	log.Printf("Environment variable MY_VAR: %s", envVar)
	log.Printf("Server is ready to handle requests at http://localhost%s", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}