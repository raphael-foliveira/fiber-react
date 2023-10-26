package main

import (
	"github.com/raphael-foliveira/fiber-react/backend/internal/api"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/db"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}
	if err := api.Start(db); err != nil {
		panic(err)
	}
}
