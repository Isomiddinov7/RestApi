package models

import (
	author "github.com/Isomiddinov7/RestApi/modules/author/models"
	book "github.com/Isomiddinov7/RestApi/modules/book/models"
)

type BookAuthor struct {
	BookAuthorId int `json:"book_author_id"`
	BookId       book.Book
	AuthorId     author.Author
}

type PostBookAuthor struct {
	BookId   int `json:"book_id"`
	AuthorId int `json:"author_id"`
}

type DeleteBookAuthor struct {
	BookAuthorId int `json:"book_author_id"`
}

type PutBookAuthor struct {
	BookAuthorId int `json:"book_author_id"`
	BookId       int `json:"book_id"`
	AuthorId     int `json:"author_id"`
}
