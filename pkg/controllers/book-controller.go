package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MuneerZaatreh/go-bookstore/pkg/models"
	"github.com/MuneerZaatreh/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBook := models.GetAllBooks()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Erorr while parsing")
	}
	BookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(BookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
   createBook := &models.Book{}
   utils.ParseBody(r,createBook)
   b:= createBook.CreateBook()
   res, _ := json.Marshal(b)
   w.Header().Set("Content-Type", "pkglication/json")
   w.WriteHeader(http.StatusOK)
   w.Write(res)

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Erorr while parsing")
	}
   book := models.DeleteBook(Id)
   res, _ := json.Marshal(book)
   w.Header().Set("Content-Type", "pkglication/json")
   w.WriteHeader(http.StatusOK)
   w.Write(res)
}
func UpdateBook (w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	Bookid := vars["id"]
	ID , err := strconv.ParseInt(Bookid,0,0)
	if err != nil {
		fmt.Println("Erorr while parsing")
	}
	Bookdetails , db:= models.GetBookById(ID)
	if updateBook.Name != ""{
		Bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		Bookdetails.Author = updateBook.Author
	}
	db.Save(&Bookdetails)
	res, _ := json.Marshal(Bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
