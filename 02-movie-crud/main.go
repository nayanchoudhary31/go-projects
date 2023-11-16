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

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	//Set the json content type
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie) // get the movie details from the request
	movie.Id = strconv.Itoa(rand.Intn(100000)) // create a new random number for movie id
	movies = append(movies, movie)             // append the movie to the list of movies
	json.NewEncoder(w).Encode(movie)           // encode the response and encode the movie
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the json content type
	w.Header().Set("Content-Type", "application/json")
	// get the id from the params
	params := mux.Vars(r)
	// loop for the movies slice and find the id in the slice
	for index, item := range movies {
		if item.Id == params["id"] {
			var movie Movie
			// delete the id from the slice
			movies = append(movies[:index], movies[index+1:]...)
			// get the details from the params
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)   // append the movie to the slice
			json.NewEncoder(w).Encode(movie) // return the json representation
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "43877", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at the port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
