package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firsname string `json:"firstname"`
	Lastname string `json:"lasname"`
}

var movies []Movie

// Get all Movies Function 
func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Delete Movie Function 
func deleteMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for index , item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index] , movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Get One Movie Using its ID

func getMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r);

	for _ , item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//Create a Movie 

func createMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r *http.Request)  {
	
	//set the json content type 
	w.Header().Set("Content-Type" , "application/json")
	// params
	params := mux.Vars(r)
	// loop over the movies , range
	// delete the movie with i.d that you've sent 
	//add the new movie - the movie that we send in the body of postman 

	for index , item := range movies{
		if item.ID == params["id"] { 
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder	(w).Encode(movie)
			return
		}
	}
}


func main()  {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1" , Isbn: "438227" , Title: "Movie One" , Director: &Director{Firsname: "Onkar" , Lastname: "Shelke"}})
	movies = append(movies, Movie{ID: "2" , Isbn: "438228" , Title: "Movie Two" , Director: &Director{Firsname: "Rohit" , Lastname: "Shelke"}})


	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("/movies" , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}" , deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at Port 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))

}