package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//creating a struct of type Movie
type Movies struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

//creating a struct of type Director
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movies

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
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

func addNewMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	json.NewDecoder(r.Body).Decode(&movie)
	random_number := strconv.Itoa(rand.Intn(100000))
	movie.Id = random_number
	movies = append(movies, movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if params["id"] == item.Id {
			var movie Movies
			json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, movie)
			break
		}
	}

	json.NewEncoder(w).Encode(&movies)

}

func main() {
	router := mux.NewRouter()
	movies = append(movies, Movies{Id: "001", Title: "KGF-1", Isbn: "100", Director: &Director{FirstName: "Prashanth", LastName: "Neel"}})
	movies = append(movies, Movies{Id: "002", Title: "Pushpa", Isbn: "101", Director: &Director{FirstName: "Sid", LastName: "Sriram"}})
	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", deleteMovieById).Methods("DELETE")
	router.HandleFunc("/movies", addNewMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	fmt.Println("Starting the server at 8080........")
	log.Fatal(http.ListenAndServe(":8080", router))
}
