create table product(
	ID serial primary key not null,
	name varchar(300) not null,
	company varchar(300) not null,
	price numeric(20, 2),
	inserted_at timestamp not null default current_timestamp
)
