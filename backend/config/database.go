package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("Mengoneksikan ke database...")
	dsn := "root:@tcp(localhost:3306)/go_rfid"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db
	fmt.Println("Koneksi database berhasil!")
}
