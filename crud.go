package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Todos []Todos `json:"todos"`
}
type Todos struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(map[string]string{"message": "hello world"})
}
func UserHandler(rw http.ResponseWriter, r *http.Request) {
	user := User{
		"0",
		"user-exemple",
		[]Todos{
			{"exemple-1", false},
			{"exemple-2", true},
			{"exemple-3", false},
		},
	}
	json.NewEncoder(rw).Encode(user)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/user/exemple", UserHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
