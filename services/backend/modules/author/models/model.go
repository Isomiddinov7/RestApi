package models

type Author struct {
	AuthorId  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Street    string `json:"street"`
	City      string `json:"city"`
}

type PostAuthor struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Street    string `json:"street"`
	City      string `json:"city"`
}

type DeleteAuthor struct {
	AuthorId int `json:"author_id"`
}

type UpdateAuthor struct {
	AuthorId  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Street    string `json:"street"`
	City      string `json:"city"`
}
