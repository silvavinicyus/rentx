package database

import (
	"database/sql"
	"rentx/src/utils/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, dbError := sql.Open("mysql", config.StrConn)

	if dbError != nil {
		return nil, dbError
	}

	if dbError = db.Ping(); dbError != nil {
		db.Close()
		return nil, dbError
	}

	return db, nil
}
