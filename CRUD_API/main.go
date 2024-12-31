package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
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



func main()  {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1" , Isbn: "438227" , Title: "Movie One" , Director: &Director{Firsname: "Onkar" , Lastname: "Shelke"}})
	movies = append(movies, Movie{ID: "2" , Isbn: "438228" , Title: "Movie Two" , Director: &Director{Firsname: "Rohit" , Lastname: "Shelke"}})


	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("/movies" , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}" , deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}