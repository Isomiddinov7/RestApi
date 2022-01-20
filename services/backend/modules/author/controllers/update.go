package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/author/models"
)

func Put(body models.UpdateAuthor) models.Author {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var author models.Author

	err = db.QueryRow(
		models.PUT_AUTHOR,
		body.AuthorId,
		body.FirstName,
		body.LastName,
		body.Street,
		body.City,
	).Scan(
		&author.AuthorId,
		&author.FirstName,
		&author.LastName,
		&author.Street,
		&author.City,
	)

	if err != nil {
		panic(err)
	}

	return author
}
