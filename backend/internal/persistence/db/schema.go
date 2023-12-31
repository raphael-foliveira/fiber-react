package db

var schema string = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS refreshtokens (
		id SERIAL PRIMARY KEY,
		token VARCHAR(255) NOT NULL,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		completed_at TIMESTAMP
	);
`
