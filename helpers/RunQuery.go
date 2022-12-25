package helpers

import (
	"database/sql"
)

func RunQuery(query string, v ...any) (*sql.Rows, error) {
	db := ConnectToDB()
	var err error
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(v...)
	return rows, err
}
