package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connection_string := "user=service-pc-api host=localhost password=pcapi sslmode=disable dbname=product-catalog"
	var err error
	db, err = sql.Open("postgres", connection_string)
	if err != nil {
		fmt.Println("sqlOpenError:", err)
	}
}

func ConnectToDB() *sql.DB {
	return db
}
