package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return db, err
	}
	_, err = db.Exec(schema)
	return db, err
}
