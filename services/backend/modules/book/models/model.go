package models

type Book struct {
	BookId          int    `json:"book_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	LongDescription string `json:"long_description"`
	CreatedAt       string `json:"created_at"`
}

type PostBook struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	LongDescription string `json:"long_description"`
}

type DeleteBook struct {
	BookId int `json:"book_id"`
}

type PutBook struct {
	BookId          int    `json:"book_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	LongDescription string `json:"long_description"`
}
