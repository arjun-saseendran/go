package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() {

	var err error
	DB, err = pgx.Connect(context.Background(), Config.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Error connectiong to DB!")
		os.Exit(1)
	}

	log.Printf("Connected with postgres database!")
}
