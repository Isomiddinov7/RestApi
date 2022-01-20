package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/customer/models"
)

func Delete(body models.DeleteCustomer) models.Customers {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var customer models.Customers

	err = db.QueryRow(
		models.DELETE_CUSTOMER,
		body.CustomerId,
	).Scan(
		&customer.CustomerId,
		&customer.FirstName,
		&customer.LastName,
		&customer.Street,
		&customer.City,
	)

	if err != nil {
		panic(err)
	}

	return customer

}
