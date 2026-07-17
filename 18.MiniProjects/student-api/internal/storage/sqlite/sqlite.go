package sqlite

import (
	"database/sql"

	"github.com/Bilal-Ahmed4/student-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS student(
	id	INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(200),
	email VARCHAR(200),
	age  INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}
