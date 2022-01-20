package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/book_comments/models"
)

func Get() []models.BookComments {
	
	db, err := sql.Open("postgres", DB.DB_CONFIG)
	
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query(models.GET_ALL_BOOK_COMMENTS)

	defer rows.Close()

	var book_comments []models.BookComments

	for rows.Next() {

		var book_comment models.BookComments

		rows.Scan(
			&book_comment.BookCommentId,
			&book_comment.Comment,
			&book_comment.SendAt,
			
			&book_comment.BookId.BookId,
			&book_comment.BookId.Title,
			&book_comment.BookId.Description,
			&book_comment.BookId.LongDescription,
			&book_comment.BookId.CreatedAt,
			
			&book_comment.CustomerId.CustomerId,
			&book_comment.CustomerId.FirstName,
			&book_comment.CustomerId.LastName,
			&book_comment.CustomerId.Street,
			&book_comment.CustomerId.City,
			
			&book_comment.BookRatin,
		)
		
		book_comments = append(book_comments, book_comment)
	}

	return book_comments
}
