package models

const GET_ALL_CUSTOMERS = `
	select
		c.customer_id,
		c.first_name,
		c.last_name,
		c.street,
		c.city
	from customers as c
`

const POST_CUSTOMER = `
	insert into customers(first_name, last_name, street, city) values($1, $2, $3, $4)
	returning
		customer_id,
		first_name,
		last_name,
		street,
		city
`

const DELETE_CUSTOMER = `
	delete from customers as c where c.customer_id = $1
	returning
		c.customer_id,
		c.first_name,
		c.last_name,
		c.street,
		c.city
`
const PUT_CUSTOMER = `
	update customers as c
		set
			first_name = case when $2='' then first_name else $2 end,
			last_name = case when $3='' then last_name else $3 end,
			street = case when $4='' then street else $4 end,
			city = case when $5='' then street else $5 end
	where c.customer_id = $1
		returning
			c.customer_id,
			c.first_name,
			c.last_name,
			c.street,
			c.city
`
