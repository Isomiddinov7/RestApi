package models

const GET_ALL_BOOKS = `
	select 
		b.book_id,
		b.title,
		b.description,
		b.long_description,
		b.created_at
	from 
		books as b
`

const POST_BOOK = `
	insert into books(title,description,long_description) values($1, $2, $3)
	returning
		book_id,
		title,
		description,
		long_description
`

const DELETE_ONE_BOOK = `
	delete from books as b where b.book_id = $1
	returning
		book_id,
		title,
		description,
		long_description
`

const PUT_BOOK = `
	update books as b
	set
		title = case when $2 = '' then title else $2 end,
		description = case when $3 = '' then description else $3 end,
		long_description = case when $4 = '' then long_description else $4 end
	where b.book_id = $1
	returning
		book_id,
		title,
		description,
		long_description
`
