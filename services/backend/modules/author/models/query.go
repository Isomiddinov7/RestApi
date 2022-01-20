package models

const GET_ALL_AUTHORS = `
	select
		a.author_id,
		a.first_name,
		a.last_name,
		a.street,
		a.city
	from 
		authors as a
`
const POST_AUTHOR = `
	insert into authors(first_name, last_name, street, city) values($1, $2, $3, $4)
	returning
		author_id,
		first_name,
		last_name,
		street,
		city
`

const DELETE_ONE_AUTHOR = `
	delete from authors as a where a.author_id = $1
	returning
		author_id,
		first_name,
		last_name,
		street,
		city
`

const PUT_AUTHOR = `
	update authors as a
	set 
		first_name =  case when $2= '' then first_name else $2 end,
		last_name = case when $3= '' then last_name else $3 end,
		street = case when $4= '' then street else $4 end,
		city = case when $5= '' then city else $5 end
	where a.author_id = $1
	returning
		author_id,
		first_name,
		last_name,
		street,
		city
`
