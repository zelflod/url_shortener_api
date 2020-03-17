package db

import (
	"database/sql"
	"io/ioutil"
)

func InitTables(db *sql.DB) error {
	file, err := ioutil.ReadFile("./sql/init.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(file))
	if err != nil {
		return err
	}

	return nil
}
