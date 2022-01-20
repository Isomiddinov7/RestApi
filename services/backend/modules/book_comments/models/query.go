package models

const GET_ALL_BOOK_COMMENTS = `
	select
		bc.book_comment_id,
		bc.comment,
		bc.send_at,
		
		b.book_id,
		b.title,
		b.description,
		b.long_description,
		b.created_at,
		
		c.customer_id,
		c.first_name,
		c.last_name,
		c.street,
		c.city,
		
		bc.book_ratin
	from
		book_comments as bc
	join books as b on bc.book_id=b.book_id
	join customers as c on bc.customer_id=c.customer_id
`
const POST_BOOK_COMMENTS = `
	insert into book_comments(comment, book_id, customer_id, book_ratin) values($1, $2, $3, $4)
	returning
		book_comment_id,
		comment,
		send_at,
		book_id,
		customer_id,
		book_ratin
`

const DELETE_BOOK_COMMENT = `
	delete from book_comments as bc where bc.book_comment_id = $1
	returning
		book_comment_id,
		comment,
		send_at,
		book_id,
		customer_id,
		book_ratin
`

const PUT_BOOK_COMMENT = `
	update book_comments as bc
	set
		comment = case when $2 = '' then comment else $2 end,
		book_id = case when $3 = 0 then book_id else $3 end,
		customer_id = case when $4 = 0 then customer_id else $4 end,
		book_ratin = case when $5 = 0 then book_ratin else $5 end
	where bc.book_comment_id = $1
	returning
		book_comment_id,
		comment,
		send_at,
		book_id,
		customer_id,
		book_ratin
`
