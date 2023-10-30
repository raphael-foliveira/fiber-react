package main

import (
	"os"

	"github.com/raphael-foliveira/fiber-react/backend/internal/api"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/db"
)

// @title Todo API
// @version 1.0
// @BasePath /api
func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := db.Connect(databaseUrl)
	if err != nil {
		panic(err)
	}
	if err := api.Start(db); err != nil {
		panic(err)
	}
}
