// Command greeter runs the greeter HTTP service.
package main

import (
	"log"
	"net/http"
	"os"

	"greeter/internal/greeting"
)

const defaultPort = "9090"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", greeting.Handler)

	log.Printf("greeter listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
