package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request){
	json.NewEncoder(rw).Encode(map[string]string{"message":"hello world"})
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/",HomeHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
