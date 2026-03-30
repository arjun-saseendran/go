package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost dbname=postgres")
	if err != nil {
		log.Fatal("Falied to connect data base!", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}
	log.Println("Succefully connected to postgres database!")
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()

}

func createTable() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY, name VARCHAR(20) NOT NULL, age INT NOT NULL, place VARCHAR(20) NOT NULL)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Couldn't create users table!")
	}
}
