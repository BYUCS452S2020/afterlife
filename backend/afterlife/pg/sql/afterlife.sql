drop table if exists tokens;
drop table if exists events;
drop table if exists users;

create table users (
	id serial PRIMARY KEY,
	email text UNIQUE NOT NULL,
	password text NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
	created_on timestamp NOT NULL,
	last_login timestamp,
	verified_alive timestamp NOT NULL
);

create table tokens (
	id serial PRIMARY KEY,
	user_id integer NOT NULL
		REFERENCES users(id)
		ON UPDATE CASCADE
		ON DELETE CASCADE,
	token text UNIQUE NOT NULL,
	created_on timestamp NOT NULL
);

create table events (
	id serial PRIMARY KEY,
	user_id integer NOT NULL
		REFERENCES users(id)
		ON UPDATE CASCADE
		ON DELETE CASCADE,
	name text NOT NULL,
	at timestamptz,
	event_type varchar(30) NOT NULL,
	email_to text[],
	email_subject text,
	email_body text
);
