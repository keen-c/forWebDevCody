package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


var store map[string]map[string]string = map[string]map[string]string{}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /key/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var value map[string]string
		if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "error decoding", http.StatusInternalServerError)
			return
		}
		store[id] = value
		fmt.Println(store)
	})

	router.HandleFunc("POST /key/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if err := json.NewEncoder(w).Encode(store[id]); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "error decoding", http.StatusInternalServerError)
			return
		}
	})

	log.Println("listening on http://localhost:9292")
	if err := http.ListenAndServe(":9292", router); err != nil {
		log.Printf("Can't listen on err %s", err)
	}
}
