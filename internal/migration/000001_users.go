package migration

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTables, downCreateTables)
}

func upCreateTables(tx *sql.Tx) error {
	query := `CREATE TABLE  users (
		id serial PRIMARY KEY,
		name varchar(255) not null,
		username varchar(255) not null unique,
		email varchar(255) not null unique,
		password_hash varchar(255) not null,
		deleted boolean DEFAULT false,
		time_deleted DATE default TO_DATE('1111111','YYYYMMDD'));`
	_, err := tx.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create users table: %v", err)
	}

	return nil
}

func downCreateTables(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		return fmt.Errorf("could not drop users table: %v", err)
	}

	return nil
}
