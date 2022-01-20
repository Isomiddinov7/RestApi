package controllers

import (
	"database/sql"

	DB "github.com/Isomiddinov7/RestApi/config"

	"github.com/Isomiddinov7/RestApi/modules/customer/models"
)

func Get() []models.Customers {

	db, err := sql.Open("postgres", DB.DB_CONFIG)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query(models.GET_ALL_CUSTOMERS)

	defer rows.Close()

	var customers []models.Customers

	for rows.Next() {

		var customer models.Customers

		rows.Scan(
			&customer.CustomerId,
			&customer.FirstName,
			&customer.LastName,
			&customer.Street,
			&customer.City,
		)

		customers = append(customers, customer)
	}

	return customers

}
