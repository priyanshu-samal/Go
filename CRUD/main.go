package main

import (
	"fmt"
	"log"
	"encoding/json"
	
	"net/http"
	
	"github.com/gorilla/mux"
)

type Movie struct{
    ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
	
	 
}

type Director struct{
    firstname string `json:"firstname"`
	lastname string `json:"lastname"`
    
}

var movies []Movie

func getMovies( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies{
		if movie.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies{
		if movie.ID == params["id"]{
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovies( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies{
		if movie.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main(){
	r:=mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{firstname: "John", lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "Movie Two", Director: &Director{firstname: "Steve", lastname: "Smith"}})





	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) 

}