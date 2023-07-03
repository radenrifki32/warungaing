package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionDb() *sql.DB {
	db, err := sql.Open("mysql", "root:rifkiganteng@tcp(localhost:3306)/warungmemie?parseTime=true")
	if err != nil {
		panic(err)
	}
	fmt.Println("connnected")
	db.SetConnMaxLifetime(10 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db
}
