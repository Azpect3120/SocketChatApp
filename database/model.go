package database

import "database/sql"

type Database struct {
	connectionString string
	database         *sql.DB
};
