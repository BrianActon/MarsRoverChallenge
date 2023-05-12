package main

import (
	"fmt"
	"net/http"

	r "mars-rover-challenge/rovers"
)

func main() {
	fmt.Println("Mars Rover Challenge Start")

	handler := http.NewServeMux()

	// json format input
	handler.HandleFunc("/input", r.Input)
	// file delivery
	handler.HandleFunc("/file", r.File)

	server := &http.Server{
		Addr:    "127.0.0.1:8095",
		Handler: handler,
	}

	server.ListenAndServe()

	fmt.Println("Mars Rover Challenge End")
}