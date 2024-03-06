package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Address string = "127.0.0.1:8081"
)

var mdbClient *mongo.Client


func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola caracola"))
	})
	http.HandleFunc("POST /notes", createNote)
	log.Fatal(http.ListenAndServe(Address, nil))
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Note: %+v", note)
}

type Note struct {
	Title string
	Tags  []string
	Text  string
	Scope Scope
}

type Note2 struct {
	Title string
	Tags  []string
	Text  string
	Scope
}

type Scope struct {
	Project string
	Area    string
}
