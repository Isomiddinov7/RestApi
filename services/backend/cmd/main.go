package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	author "github.com/Isomiddinov7/RestApi/modules/author/routers"
	book "github.com/Isomiddinov7/RestApi/modules/book/routers"
	book_author "github.com/Isomiddinov7/RestApi/modules/book_author/routers"
	book_comment "github.com/Isomiddinov7/RestApi/modules/book_comments/routers"
	customer "github.com/Isomiddinov7/RestApi/modules/customer/routers"
)

func main() {

	r := mux.NewRouter()

	// BOOK
	r.HandleFunc("/book/list", book.GET).Methods("GET")
	r.HandleFunc("/book/create", book.POST).Methods("POST")
	r.HandleFunc("/book/delete", book.DELETE).Methods("DELETE")
	r.HandleFunc("/book/update", book.PUT).Methods("PUT")

	// AUTHOR
	r.HandleFunc("/author/list", author.GET).Methods("GET")
	r.HandleFunc("/author/create", author.POST).Methods("POST")
	r.HandleFunc("/author/delete", author.DELETE).Methods("DELETE")
	r.HandleFunc("/author/update", author.PUT).Methods("PUT")

	// CUSTOMER
	r.HandleFunc("/customer/list", customer.GET).Methods("GET")
	r.HandleFunc("/customer/create", customer.POST).Methods("POST")
	r.HandleFunc("/customer/delete", customer.DELETE).Methods("DELETE")
	r.HandleFunc("/customer/update", customer.PUT).Methods("PUT")

	// BOOK - AUTHOR
	r.HandleFunc("/book_author/list", book_author.GET).Methods("GET")
	r.HandleFunc("/book_author/create", book_author.POST).Methods("POST")
	r.HandleFunc("/book_author/delete", book_author.DELETE).Methods("DELETE")
	r.HandleFunc("/book_author/update", book_author.PUT).Methods("PUT")

	// BOOK - COMMENT
	r.HandleFunc("/book_comment/list", book_comment.GET).Methods("GET")
	r.HandleFunc("/book_comment/create", book_comment.POST).Methods("POST")
	r.HandleFunc("/book_comment/delete", book_comment.DELETE).Methods("DELETE")
	r.HandleFunc("/book_comment/update", book_comment.PUT).Methods("PUT")

	fmt.Println("Running server port in :8080")

	http.ListenAndServe(":8080", r)
}
