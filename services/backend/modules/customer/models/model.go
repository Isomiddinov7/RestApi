package models

type Customers struct {
	CustomerId int    `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Street     string `json:"street"`
	City       string `json:"city"`
}

type PostCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Street    string `json:"street"`
	City      string `json:"city"`
}

type DeleteCustomer struct {
	CustomerId int `json:"customer_id"`
}

type PutCustomer struct {
	CustomerId int    `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Street     string `json:"street"`
	City       string `json:"city"`
}
