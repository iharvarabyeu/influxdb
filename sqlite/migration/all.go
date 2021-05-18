package migration

import (
	"database/sql"
	"log"
	"sync"

	"github.com/lopezator/migrator"
)

type SqliteMigrator struct {
	mu sync.Mutex
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

	// Migrate up
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := m.Migrate(s.DB); err != nil {
		log.Fatal(err)
	}
}
