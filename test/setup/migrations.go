package setup

import (
	"database/sql"
)

func RunMigrations(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE users (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
		CREATE TABLE events (
			id varchar(36) PRIMARY KEY,
			description TEXT NOT NULL,
			date timestamp NOT NULL,
			address TEXT NOT NULL,
			user_id varchar(36) NOT NULL
		);
	`)
	return err
}
