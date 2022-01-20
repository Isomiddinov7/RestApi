drop table if exists books cascade;
drop table if exists authors cascade;
drop table if exists customers cascade;
drop table if exists book_author cascade;
drop table if exists book_comments cascade;

create table books(
    book_id serial not null primary key,
    title varchar(128) not null,
    description varchar(256) not null,
    long_description text not null,
    created_at timestamp with time zone default current_timestamp
);

create table authors(
    author_id serial not null primary key,
    first_name varchar(128) not null,
    last_name varchar(128) not null,
    street varchar(64) not null,
    city varchar(128) not null
);

create table customers(
    customer_id serial not null primary key,
    first_name varchar(128) not null,
    last_name varchar(128) not null,
    street varchar(64) not null,
    city varchar(128) not null
);

create table book_author(
    book_author_id serial not null primary key,
    book_id int not null references books(book_id),
    author_id int not null references authors(author_id)
);

create table book_comments(
    book_comment_id serial not null primary key,
    comment text not null,
    send_at timestamp with time zone default current_timestamp,
    book_id int not null references books(book_id),
    customer_id int not null references customers(customer_id),
    book_ratin smallint not null
);