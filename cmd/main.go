package main

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("starting server ...")

	r := mux.NewRouter()
	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/{name}/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

type Res struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const base32alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func ID() string {
	src := make([]byte, 26)
	rand.Read(src)

	for i := range src {
		src[i] = base32alphabet[src[i]%32]
	}
	return string(src)
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := ID()
	id = id[:8]

	vars := mux.Vars(r)
	name := vars["name"]

	res := Res{
		ID:   id,
		Name: name,
	}

	encoder := json.NewEncoder(w)

	err := encoder.Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
