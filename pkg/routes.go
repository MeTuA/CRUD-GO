package pkg

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//defining routes
func Routes() {
	routes := mux.NewRouter()
	s := routes.PathPrefix("/api").Subrouter()

	s.HandleFunc("/createUser", CreateUser).Methods("POST")
	s.HandleFunc("/getUser", GetUser).Methods("GET")
	s.HandleFunc("/updateUser", UpdateUser).Methods("POST")
	s.HandleFunc("/deleteUser/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", s))
}
