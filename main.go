package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movies struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json: "director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movies

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	fmt.Println("Inside get All movies")
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	movies = append(movies, Movies{Id: "001", Title: "KGF-1", Isbn: "100", Director: &Director{FirstName: "Prashanth", LastName: "Neel"}})
	movies = append(movies, Movies{Id: "002", Title: "Pushpa", Isbn: "101", Director: &Director{FirstName: "Sid", LastName: "Sriram"}})
	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	fmt.Println("Starting the server at 8080........")
	log.Fatal(http.ListenAndServe(":8080", router))
}
