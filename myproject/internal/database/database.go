package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("postgres", "postgres://user:password@localhost/database_name?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func saveOrder(order Order) error {

	_, err := db.Exec("INSERT INTO orders (user_id, items) VALUES ($1, $2)", order.UserID, order.Items)
	return err
}
