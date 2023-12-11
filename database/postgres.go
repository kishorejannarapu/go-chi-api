package database

import (
	"fmt"
	models "go-chi-api/model"

	"github.com/go-pg/pg"
)

func ConnectToPostgres(dbURL string) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     dbURL,
		User:     "username",
		Password: "password",
		Database: "database_name",
	})

	if db == nil {
		return nil, fmt.Errorf("failed to connect to database")
	}

	return db, nil
}

func CreateUser(db *pg.DB, user *models.User) error {
	// Implement logic to create user in database
	return nil
}

// Implement other database operations
