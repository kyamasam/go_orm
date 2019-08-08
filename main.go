package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World")
}
var port="8080"

func handleRequests()  {
	//define routes
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", GetUser).Methods("GET")
	myRouter.HandleFunc("/users", NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

//define main entry

func main()  {
	fmt.Printf("Starting Sever on http://127.0.0.1:%s",port )
	//call the migration to create the database tables
	InitialMigration()
	//start the server
	handleRequests()
}