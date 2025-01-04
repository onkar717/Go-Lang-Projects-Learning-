package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/onkar717/pkg/utils"
	"github.com/onkar717/pkg/models"
)

var NewBook models.Book 

func GetBook(w http.ResponseWriter, r *http.Request)  {
	newBooks := models.GetAllBooks()
	res , _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type" ,"applicatopm/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request)  {
	CreateBook := &models.Book{}
	utils.ParseBody(r , CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId , 0 , 0)
	if err != nil {
		fmt.Println("Error While Deleting ")
	}
	book := models.Delete(ID)
	res , _ := json.Marshal(book)
	w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request)  {
	var updateBook = &models.Book{}
	utils.ParseBody(r , updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId , 0 , 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	booksDetails , db := models.GetBookById(ID)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name;
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author;
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication;
	}

	db.Save(&booksDetails)
	res , _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}