package main

import (
	"log"
	"net/http"

	"./dummydb"
	"./endpoint"
	"./server"
)

const serverAddress = ":8080"

func main() {
	dummydb.InitDB()

	epr := new(endpoint.Request)
	mux := http.NewServeMux()
	mux.HandleFunc("/", epr.Parse)

	srv := server.New(mux, serverAddress)

	// In the real world we would use ListenAndServeTLS() with certificate details
	// so that our TLS settings are actually applied.
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
