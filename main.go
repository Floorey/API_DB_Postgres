package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	//p -> postgres
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("Environment variable DB_PASSWORD not set")
	}

	connStr := fmt.Sprintf("user=postgres password=%s dbname=users_go sslmode=disable", password)

	// connection to PostgresSQL-DB

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Ping the database to ensure the connection was successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to PostgresSQL database established successfully!")

}
