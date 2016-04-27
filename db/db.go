package db

import (
	"database/sql"
	"log"

	"github.com/ReneVallecillo/OTBackend/config"
)

//NewDB checks for a db, if not creates a new one.
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, table := range config.InitSqls {
		_, err := db.Exec(table)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return db, nil

}
