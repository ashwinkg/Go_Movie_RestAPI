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
	// w.WriteHeader(http.StatusOK)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if params["id"] == item.Id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if params["id"] == item.Id {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			break
		}
	}

}

func main() {
	router := mux.NewRouter()
	movies = append(movies, Movies{Id: "001", Title: "KGF-1", Isbn: "100", Director: &Director{FirstName: "Prashanth", LastName: "Neel"}})
	movies = append(movies, Movies{Id: "002", Title: "Pushpa", Isbn: "101", Director: &Director{FirstName: "Sid", LastName: "Sriram"}})
	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", deleteMovieById).Methods("DELETE")
	fmt.Println("Starting the server at 8080........")
	log.Fatal(http.ListenAndServe(":8080", router))
}
