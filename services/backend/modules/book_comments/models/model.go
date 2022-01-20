package models

import (
	book "github.com/Isomiddinov7/RestApi/modules/book/models"
	customer "github.com/Isomiddinov7/RestApi/modules/customer/models"
)

type BookComments struct {
	BookCommentId int    `json:"book_comment_id"`
	Comment       string `json:"comment"`
	SendAt        string `json:"send_at"`
	BookId        book.Book
	CustomerId    customer.Customers
	BookRatin     int `json:"book_ratin"`
}

type PostBookComments struct {
	Comment    string `json:"comment"`
	SendAt     string `json:"send_at"`
	BookId     int    `json:"book_id"`
	CustomerId int    `json:"customer_id"`
	BookRatin  int    `json:"book_ratin"`
}

type DeleteBookComments struct {
	BookCommentId int `json:"book_comment_id"`
}

type PutBookComments struct {
	BookCommentId int    `json:"book_comment_id"`
	Comment       string `json:"comment"`
	BookId        int    `json:"book_id"`
	CustomerId    int    `json:"customer_id"`
	BookRatin     int    `json:"book_ratin"`
}
