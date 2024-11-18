package main

import (
	person "TestingGo/cmd/internal"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/people", getPeople)
	http.HandleFunc("/new-person", newPerson)
	http.HandleFunc("/delete-person", deletePerson)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
	log.Println("Server is running on port 8080")
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(person.People)
	if err != nil {
		http.Error(w, "Failed to encode people", http.StatusInternalServerError)
	}
}

func newPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}

	person.New(name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(newPerson)
	if err != nil {
		http.Error(w, "Failed to encode new person", http.StatusInternalServerError)
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
		return
	}

	for i, p := range person.People {
		if p.Id == id {
			person.People = append(person.People[:i], person.People[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}
