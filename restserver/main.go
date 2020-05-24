package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Start is a method that starts the server.
func main() {
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/putfile", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(100000000)
		fmt.Printf("request received: %v - %v\n", r.Body, r.MultipartForm)
		w.WriteHeader(200)
	}).Methods("POST", "OPTIONS")

	http.Handle("/", accessControl(mux))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// CORS control
func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// func getMongoConnection(connectionString string) *mongo.Client
