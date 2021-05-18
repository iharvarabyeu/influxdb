package migration

import (
	"database/sql"
	"log"

	"github.com/lopezator/migrator"
)

// Up runs the migrations
func Up(db *sql.DB) {
	// Configure migrations
	m, err := migrator.New(
		migrator.Migrations(
			&migrator.Migration{
				Name: "Create table foo",
				Func: func(tx *sql.Tx) error {
					if _, err := tx.Exec("CREATE TABLE foo (id INT PRIMARY KEY)"); err != nil {
						return err
					}
					return nil
				},
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate up
	if err := m.Migrate(db); err != nil {
		log.Fatal(err)
	}
}
