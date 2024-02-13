package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type Store struct {
	Name    string `json:"name,omitempty"`
	Youtube string `json:"youtube,omitempty"`
}

var store map[string]map[string]string = make(map[string]map[string]string)
var anotherstore map[string]Store
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
    var anotherstorevalue Store
		if err := json.NewDecoder(r.Body).Decode(&anotherstorevalue); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "error decoding", http.StatusInternalServerError)
			return
		}

		store[id] = value
    anotherstore[id] = anotherstorevalue
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
		log.Printf("Can't listen on err %v", err)
	}
}
