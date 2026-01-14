package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	message := r.URL.Query().Get("message")
	if message == "" {
		http.Error(w, "missing message", http.StatusBadRequest)
		return
	}

	log.Printf("received message: %s", message)
	fmt.Fprintf(w, "ok: %s\n", message)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/message", handleMessage)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
