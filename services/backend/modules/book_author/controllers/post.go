package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/book_author/models"
)

func Post(body models.PostBookAuthor) models.BookAuthor {
	
	db, err := sql.Open("postgres", DB.DB_CONFIG)
	
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var book_author models.BookAuthor

	err = db.QueryRow(
		models.POST_BOOK_AUTHOR,
		body.BookId,
		body.AuthorId,
	).Scan(
		&book_author.BookAuthorId,
		&book_author.BookId.BookId,
		&book_author.AuthorId.AuthorId,
	)

	if err != nil {
		panic(err)
	}

	return book_author
}
