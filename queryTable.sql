create table purchases (
	id BIGSERIAL primary key,
	product_id int8 notnull,
	user_id int8 notnull,
	qtty int notnull,
	receive int,
	status varchar(50),
	created_at timestamp(6),
	updated_at timestamp(6)
)

create table carts (
	id BIGSERIAL primary key,
	product_id int8 notnull,
	user_id int8 notnull,
	qtty int notnull,
	Amount int8,
	created_at timestamp(6),
	updated_at timestamp(6)
)

create table histories (
	id BIGSERIAL primary key,
	product_id int8 notnull,
	user_id int8 notnull,
	qtty int notnull,
	amount int8,
	status varchar(50),
	created_at timestamp(6),
	updated_at timestamp(6)
)

create table products (
	id BIGSERIAL primary key,
	name varchar(100) notnull,
	detail varchar(100) notnull,
	price int8 notnull,
    stock int8,
	user_id int8,
	created_at timestamp(6),
	updated_at timestamp(6)
)

create table users (
	id BIGSERIAL primary key,
	name varchar(100) notnull,
	email varchar(100) unique notnull,
    password varchar(100) notnull,
	created_at timestamp(6),
	updated_at timestamp(6)
)