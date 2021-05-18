package migration

import (
	"database/sql"
	"log"

	"github.com/lopezator/migrator"
)

type SqliteMigrator struct {
	DB *sql.DB
}

// Up runs the migrations
func (s *SqliteMigrator) Up() {
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

	if err := m.Migrate(s.DB); err != nil {
		log.Fatal(err)
	}
}
