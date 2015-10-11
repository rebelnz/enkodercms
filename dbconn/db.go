package dbconn

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// connection parameters get search/replaced by enkdr-script
func NewDB() *sqlx.DB {
	s := fmt.Sprintf("user=%s password=%s dbname=%s port=5432 sslmode=disable", "enkoder", "iediexoh", "enkoder")
	db, err := sqlx.Open("postgres", s)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Could Not connect")
	}
	return db
}
