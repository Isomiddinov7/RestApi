package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"
	
	"github.com/Isomiddinov7/RestApi/modules/book/models"
)

func Delete(body models.DeleteBook) models.Book {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var book models.Book

	err = db.QueryRow(
		models.DELETE_ONE_BOOK,
		body.BookId,
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
