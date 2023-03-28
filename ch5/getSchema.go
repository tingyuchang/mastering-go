package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) != 6 {
		fmt.Println("Please provide: Hostname | Port | Username | Password | DB ")
		return
	}

	host := arguments[1]
	portStr := arguments[2]
	username := arguments[3]
	pasword := arguments[4]
	database := arguments[5]

	port, err := strconv.Atoi(portStr)

	if err != nil {
		fmt.Println("Port is invalid ", err)
		return
	}

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, pasword, database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("OpenDB: ", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "datname" FROM "pg_database" WHERE datistemplate=false`)
	if err != nil {
		fmt.Println("Query: ", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan: ", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()

	query := `SELECT table_name FROM information_schema.tables WHERE table_schema= 'public' ORDER BY table_name`
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println("Query", query, " ", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan: ", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
