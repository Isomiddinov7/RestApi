package models

const GET_ALL_BOOK_AUTHOR = `
	select
		ba.book_author_id,
		b.book_id,
		b.title,
		b.description,
		b.long_description,
		b.created_at,
		a.author_id,
		a.first_name,
		a.last_name,
		a.street,
		a.city
	from
		book_author as ba
	join books as b on ba.book_id=b.book_id
	join authors as a on ba.author_id=a.author_id
`

const POST_BOOK_AUTHOR = `
	insert into book_author(book_id, author_id) values($1, $2)
	returning
		book_author_id,
		book_id,
		author_id
`

const DELETE_BOOK_AUTHOR = `
	delete from book_author as b where b.book_author_id = $1
	returning
		book_author_id,
		book_id,
		author_id
`

const PUT_BOOK_AUTHOR = `
	update book_author as ba
	set
		book_id = case when $2 = 0 then book_id else $2 end,
		author_id = case when $3 = 0 then author_id else $3 end
	where ba.book_author_id = $1
	returning
		book_author_id,
		book_id,
		author_id
`
