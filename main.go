package main

import (
	"log"
	"net/http"
)

const addr = "0.0.0.0:8080"

func main() {
	mux := http.NewServeMux()

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}
