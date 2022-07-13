package main

import (
	// Built in packages
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	// External packages
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// Set random movies
func randMovies() {
	movies = append(movies, Movie{ID: "1", Isbn: "213416", Title: "Random Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "813412", Title: "Second Random Movie", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "3", Isbn: "216415", Title: "A Random Three Movie", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
}

// Get all the movies handler
func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

// Get a single movies with the movie id
func getAMovie(res http.ResponseWriter, req *http.Request) {
	// Getting all the parameters from the URI
	params := mux.Vars(req)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			// If the movie is available then send a response and break the loop
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(movie)
			return
		}
	}
	// Incase the movie was not found
	http.Error(res, "Movie not found", http.StatusNotFound)
}

// Create a movie with the right data
func createMovie(res http.ResponseWriter, req *http.Request) {
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	// New movie's ID and Isbn
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	// Send the response
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movie)
}

// Update a movie with the movie id
func updateMovie(res http.ResponseWriter, req *http.Request) {
	// Find the movie with the id
	params := mux.Vars(req)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			// Remove the old movie from the db (in this case movies slice)
			movies = append(movies[:index], movies[index+1:]...)
			// Update the selected movie with the body data
			var updatedMovie Movie
			_ = json.NewDecoder(req.Body).Decode(&updatedMovie)
			updatedMovie.ID = movie.ID
			// Insert the new movie into the db (in this case movies slice)
			movies = append(movies, updatedMovie)
			// Send a response with the new movie
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(updatedMovie)
			return
		}
	}
	// Incase the movie was not found send an Error response
	http.Error(res, "Invalid movie ID", http.StatusNotFound)
}

// Remove a movie with the movie id
func removeMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	// Getting all the parameters from the URI
	params := mux.Vars(req)
	// Loop over all the movies and match their ids.
	for index, item := range movies {
		if item.ID == params["id"] {
			// If the id matches then take the index and perform an manual item remove operation.
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func main() {
	randMovies()

	router := mux.NewRouter()
	// Routes and handlers
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getAMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", removeMovie).Methods("DELETE")

	// Start the server and catch errors
	fmt.Printf("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("Unexpected error", err)
	}
}
