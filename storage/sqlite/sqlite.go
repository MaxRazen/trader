package sqlite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Migrate() {
	//q := `CREATE TABLE orders(
	//		id INTEGER PRIMARY KEY AUTOINCREMENT,
	//		amount FLOAT NOT NULL ,
	//		buy FLOAT NOT NULL ,
	//		sell FLOAT,
	//		opened_at TIMESTAMP NOT NULL DEFAULT NOW(),
	//		closed_at TIMESTAMP,
	//	);`
	//r, err := s.db.Exec(q)
	//if err != nil {
	//
	//}

}
