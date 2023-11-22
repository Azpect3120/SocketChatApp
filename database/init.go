package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func CreateDatabase(connectionString string) *Database {
	database := &Database{
		connectionString: connectionString,
		database: nil,
	}
	return database
}

func (db *Database) Connect() error {
	conn, err := sql.Open("postgres", db.connectionString)
	db.database = conn
	return err
}
