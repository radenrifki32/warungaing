package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionDb() *sql.DB {
	db, err := sql.Open("mysql", "root:rifkiganteng@tcp(localhost:3306)/warungku?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(10 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db
}
