package main

import (
	"crud/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(map[string]string{"message": "hello world"})
}
func UserHandler(rw http.ResponseWriter, r *http.Request) {
	user := model.User{
		"0",
		"user-exemple",
		[]model.Todos{
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
	r.HandleFunc("'/user/exemple'", UserHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
