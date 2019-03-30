package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// User is sample structure.
type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func main() {
	db, err := sqlx.Connect("mysql", "root:@(localhost:43306)/sqlx_development?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := addUser(db, "John Doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}

	user := User{}
	userState := `SELECT * FROM user WHERE id = ?`
	err = db.QueryRowx(userState, id).StructScan(&user)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Added user was %v\n", user)
}

func addUser(db *sqlx.DB, n, e string) (int64, error) {
	userState := `INSERT INTO user (name, email, created_at, updated_at) VALUES (?, ?, ?, ?)`
	r := db.MustExec(userState, n, e, time.Now(), time.Now())
	return r.LastInsertId()
}
