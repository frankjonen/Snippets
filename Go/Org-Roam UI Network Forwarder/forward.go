package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Create a reverse proxy to forward requests to org-roam UI
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:35901",
	})

	// Create a new HTTP server
	server := http.Server{
		Addr:    ":8080", // Port on which the Golang app will listen
		Handler: proxy,
	}

	log.Println("Starting server on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

