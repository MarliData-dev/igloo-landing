package main

import (
	"log"
	"net/http"
)

func main() {
	const addr = ":8080"

	fs := http.FileServer(http.Dir("."))
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	log.Printf("serving igloo-landing on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
