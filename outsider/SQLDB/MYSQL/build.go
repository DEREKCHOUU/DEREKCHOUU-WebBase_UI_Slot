package MYSQL

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_global "go/UI/config"

	_ "github.com/go-sql-driver/mysql"
)

func Opensql(cfg _global.DB, utc string) (*sql.DB, error) {
	dburl := fmt.Sprintf("%s:%s@tcp(%s:%v)/UI", cfg.DBac, cfg.DBpw, cfg.DBip, cfg.DBport)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", utc)
	dsn := fmt.Sprintf("%s?%s", dburl, val.Encode())
	DB, err := sql.Open("mysql", dsn)
	return DB, err
}

func Build(DB *sql.DB) {
	f, err := os.ReadFile("./sql/mysqldb/build.sql")
	if err != nil {
		log.Fatal("cannot read the sql srcipt ", err)
	}
	s := string(f)
	_, err = DB.Exec(s)
	if err != nil {
		log.Fatal("cannot build the sql DB ", err)
	}
}

// func Preparestmt() {
// 	stmt, err := DB.Prepare("SELECT * WHERE ID = ?")
// 	if err != nil {
// 		log.Fatal("query prepare fail: ", err)
// 	}
// 	err = stmt.QueryRow(12).Scan(&devOs)
// 	if err != nil {
// 		log.Fatal("query excuit fail:", err)
// 	}
// }
