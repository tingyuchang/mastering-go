package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// DataBase

var HOSTNAME = "localhost"
var DBPORT = "5432"
var DBUSER = "matt"
var DBPASS = "pass"
var DBNAME = "mydb"

func connectPostgres() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOSTNAME, DBPORT, DBUSER, DBPASS, DBNAME)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

func DeleteUser(id int) bool {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to DB")
		db.Close()
		return false
	}
	defer db.Close()

	t := FindUser(id)

	if t.Id == 0 {
		log.Println("User ", id, " does not exist.")
		return false
	}

	statement, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		log.Println(err)
		return false
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Println("Delete user failed: ", id)
		return false
	}

	return true
}

func FindUser(id int) *User {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to DB")
		db.Close()
		return nil
	}
	defer db.Close()

	statement := fmt.Sprintf(`SELECT id, username, password FROM users WHERE id='%d'`, id)
	var user User
	err := db.QueryRow(statement).Scan(&user)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &user
}

// isUserExist return true if username is not exist in database
func isUserExist(username string) bool {
	db := connectPostgres()

	if db == nil {
		fmt.Println("Cannot connect to database.")
		return false
	}
	defer db.Close()

	statement := fmt.Sprintf(`SELECT username FROM users WHERE username='%s'`, username)
	var user string
	err := db.QueryRow(statement).Scan(&user)
	if err != nil {
		log.Println(err)
		return false
	}

	if username != user {
		return true
	}
	return false
}

func ListUsers() []User {
	db := connectPostgres()

	if db == nil {
		fmt.Println("Cannot connect to database.")
		return nil
	}
	defer db.Close()

	statement := fmt.Sprintf(`SELECT id, username, password, lastlogin FROM users`)
	rows, err := db.Query(statement)

	if err != nil {
		log.Println(err)
		return nil
	}
	var all []User
	var id int
	var username string
	var password string
	var lastLogin sql.NullInt64
	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &lastLogin)
		if err != nil {
			log.Println(err)
			return nil
		}
		user := User{
			Id:        id,
			Username:  username,
			Password:  password,
			LastLogin: lastLogin.Int64,
		}

		all = append(all, user)
	}
	defer rows.Close()

	return all
}

func isUserValid(user User) bool {
	db := connectPostgres()
	if db == nil {
		log.Println("Cannot connect to database.")
		return false
	}
	defer db.Close()

	statement := fmt.Sprintf(`SELECT idm username, password FROM users WHERE username='%s'`, user.Username)
	var aUser User
	var id int
	var username string
	var password string
	err := db.QueryRow(statement).Scan(&id, &username, &password)

	if err != nil {
		log.Println(err)
		return false
	}

	aUser = User{
		Id:       id,
		Username: username,
		Password: password,
	}

	if aUser.Password == user.Password {
		return true
	}

	return false
}

func InsertUser(user User) error {
	db := connectPostgres()

	if db == nil {
		fmt.Println("Cannot connect to database.")
		return fmt.Errorf("Cannot connect to database. \n")
	}
	defer db.Close()

	// check user exist or not
	if !isUserExist(user.Username) {
		return fmt.Errorf("%s already exist in database \n", user.Username)
	}

	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES ($1, $2)")

	if err != nil {
		return err
	}

	n, err := stmt.Exec(user.Username, user.Password)

	if err != nil {
		return err
	}

	fmt.Println(n)
	return nil
}
