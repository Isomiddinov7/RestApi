package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/book/models"
)

func Update(body models.PutBook) models.Book {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var book models.Book

	err = db.QueryRow(
		models.PUT_BOOK,
		body.BookId,
		body.Title,
		body.Description,
		body.LongDescription,
	).Scan(
		&book.BookId,
		&book.Title,
		&book.Description,
		&book.LongDescription,
	)

	if err != nil {
		panic(err)
	}

	return book
}
