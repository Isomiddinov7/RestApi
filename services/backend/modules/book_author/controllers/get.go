package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/book_author/models"
)

func Get() []models.BookAuthor {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query(models.GET_ALL_BOOK_AUTHOR)

	defer rows.Close()

	var book_authors []models.BookAuthor

	for rows.Next() {

		var book_author models.BookAuthor

		rows.Scan(
			&book_author.BookAuthorId,

			&book_author.BookId.BookId,
			&book_author.BookId.Title,
			&book_author.BookId.Description,
			&book_author.BookId.LongDescription,
			&book_author.BookId.CreatedAt,
			
			&book_author.AuthorId.AuthorId,
			&book_author.AuthorId.FirstName,
			&book_author.AuthorId.LastName,
			&book_author.AuthorId.Street,
			&book_author.AuthorId.City,
		)

		book_authors = append(book_authors, book_author)
	}

	return book_authors
}
