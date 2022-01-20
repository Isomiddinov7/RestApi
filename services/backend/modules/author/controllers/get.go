package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/author/models"
)

func Get() []models.Author {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query(models.GET_ALL_AUTHORS)

	defer rows.Close()

	var authors []models.Author

	for rows.Next() {

		var author models.Author

		rows.Scan(
			&author.AuthorId,
			&author.FirstName,
			&author.LastName,
			&author.Street,
			&author.City,
		)

		authors = append(authors, author)
	}

	return authors
}
