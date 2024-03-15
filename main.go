package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	//p -> postgres
	password := "200Luckylou2010&"

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
