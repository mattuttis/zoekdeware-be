package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	birthDate, _ := time.Parse(time.DateOnly, "1976-02-07")
	router := mux.NewRouter()
	router.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		var members = []Member{
			{Id: "1", FirstName: "John", LastName: "Doe", BirthDate: birthDate},
			{Id: "2", FirstName: "Jane", LastName: "Doe", BirthDate: birthDate},
			{Id: "3", FirstName: "John", LastName: "Smith", BirthDate: birthDate},
		}
		if err := json.NewEncoder(w).Encode(members); err != nil {
			panic(err)
		}
	})

	// Handle CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}

type Member struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
}
