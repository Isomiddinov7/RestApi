package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/book_comments/models"
)

func Update(body models.PutBookComments) models.BookComments {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var book_comment models.BookComments

	err = db.QueryRow(
		models.PUT_BOOK_COMMENT,
		body.BookCommentId,
		body.Comment,
		body.BookId,
		body.CustomerId,
		body.BookRatin,
	).Scan(
		&book_comment.BookCommentId,
		&book_comment.Comment,
		&book_comment.SendAt,
		&book_comment.BookId.BookId,
		&book_comment.CustomerId.CustomerId,
		&book_comment.BookRatin,
	)

	if err != nil {
		panic(err)
	}

	return book_comment
}
