package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sawyer/go-bookstore/pkg/models"
	"github.com/sawyer/go-bookstore/pkg/utils"
	"github.com/stackpearson/go-bookstore/pkg/models"
)

//this is a var based off our Book struct from the models file
var NewBook models.Book

//need a response writer (w points to this) & a request (r points to this)
func GetBook(w http.ResponseWriter, r *http.Request) {
	//this lets us pull all books and store it inside of newBooks
	newBooks := models.GetAllBooks()
	//this converts the response into json
	res, _ := json.Marshal(newBooks)
	//just setting hetter
	w.Header().Set("Content-Type", "pkglication/json")
	//this returns http status 200 (ok)
	w.WriteHeader(http.StatusOK)
	//this sends the response from what we queried & decoded from our DB
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//getting access to our request so we can get the id
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	//book id will be a string, we're just converting it to an int here
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	//now that we have our ID extracted, we can pass it to our GetBooksByID model
	bookDetails, _ := models.GetBookById(ID)
	//now send the book details to our front end
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	//this references the CreateBook function in our book-controller
	b := CreateBook.CreateBook()
	//now we need to fom our response
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
bookId:
	vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.WriteResponse(res)
}
