package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    fmt.Printf("Starting the webserver and listening\n")
    log.Fatal(http.ListenAndServe(":8000", router))
}


func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get People API called\n")
}
func GetPerson(w http.ResponseWriter, r *http.Request) {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}