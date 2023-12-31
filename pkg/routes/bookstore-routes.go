package routes
import (
	"github.com/gorilla/mux"
	"github.com/MuneerZaatreh/go-bookstore/pkg/controllers"
)
var RegisterBookStoreRouters = func (router * mux.Router) {
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}",controllers.DeleteBook).Methods("DELETE")
}