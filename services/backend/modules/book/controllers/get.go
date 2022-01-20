package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"
	
	"github.com/Isomiddinov7/RestApi/modules/book/models"
)

func Get() []models.Book {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query(models.GET_ALL_BOOKS)

	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		rows.Scan(
			&book.BookId,
			&book.Title,
			&book.Description,
			&book.LongDescription,
			&book.CreatedAt,
		)

		books = append(books, book)
	}

	return books
}
