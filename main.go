package main

import (
	"net/http"
	"log"
)


func main(){
	const filepathRoot = "."
	const port = "8080"

	mu := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(filepathRoot))

	mu.Handle("/app/", http.StripPrefix("/app", fileServer))
	mu.HandleFunc("/healthz", endpointReadiness)

	server := http.Server{
		Addr: ":" + port,
		Handler: mu,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}


func endpointReadiness(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}