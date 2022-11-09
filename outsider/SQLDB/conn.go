package MYSQL

import "database/sql"

func Test(db *sql.DB) error {
	err := db.Ping()
	return err
}
