package psql

import (
	"database/sql"
	"log"
)

func PsqlSetup() (*sql.DB, error) {
	psqlInfo := PsqlInfo()

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error PsqlInfoSetup: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error PsqlInfoSetup: %s", err)
		return nil, err
	}
	return db, err
}
