package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Bilal-Ahmed4/student-api/internal/config"
	"github.com/Bilal-Ahmed4/student-api/internal/types"
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

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare(`INSERT INTO student (name,email,age)VALUES(?,?,?)`) //to avoid sql injection
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastid, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare(`SELECT name, email, age FROM student WHERE id = ?`)
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student
	student.Id = int64(id)

	err = stmt.QueryRow(id).Scan(&student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %d", id)
		}
		return types.Student{}, err
	}

	return student, nil

}
