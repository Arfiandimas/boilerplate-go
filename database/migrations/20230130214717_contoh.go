package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upContoh, downContoh)
}

func upContoh(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE example SET email='admin' WHERE email='root';")
	if err != nil {
		return err
	}
	return nil
}

func downContoh(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE example SET email='admin' WHERE email='root';")
	if err != nil {
		return err
	}
	return nil
}
