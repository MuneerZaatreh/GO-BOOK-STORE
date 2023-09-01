package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MuneerZaatreh/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
	fmt.Println("The Port is starting on 9010")

}
