package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thneves/go-bookstore/pkg/models"
	"github.com/thneves/go-bookstore/pkg/util"
)

var NewBook models.Book

func getBooks(write http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	write.Header().Set("Content-Type", "pkglication/json")
	write.WriteHeader(http.StatusOK)
	write.Write(res)
}

func GetBookById(write http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	// Access request, get book Id inside request
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID) // _ to not use the DB returned in the model

	res, _ := json.Marshal(bookDetails) // have to send a json response to the user
	write.Header().Set("Content-Type", "pkglication/json")
	write.WriteHeader(http.StatusOK)
	write.Write(res)
}
