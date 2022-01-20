package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/author/models"
)

func Delete(body models.DeleteAuthor) models.Author {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var author models.Author

	err = db.QueryRow(
		models.DELETE_ONE_AUTHOR,
		body.AuthorId,
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
