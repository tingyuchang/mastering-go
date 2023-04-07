package dbpost

import (
	"database/sql"
	"fmt"
	"strings"
)

/*

The package works on 2 tables on a PostgreSQL data base server.
The names of the tables are:
    * Users
    * Userdata
The definitions of the tables in the PostgreSQL server are:
    CREATE TABLE Users (
        ID SERIAL,
        Username VARCHAR(100) PRIMARY KEY”
	);
    CREATE TABLE Userdata (
        UserID Int NOT NULL,
        Name VARCHAR(100),
        Surname VARCHAR(100),
        Description VARCHAR(200)
    );
    This is rendered as code
This is not rendered as code”

*/

import (
	_ "github.com/lib/pq"
)

/*
This block of global variables holds the connection details to the Postgres server
    Hostname: is the IP or the hostname of the server
    Port: is the TCP port the DB server listens to
    Username: is the username of the database user
    Password: is the password of the database user
    Database: is the name of the Database in PostgreSQL
*/

type UserData struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

var (
	Hostname = ""
	Port     = 5432
	Username = ""
	Password = ""
	Database = ""
)

// openConnection() is for opening the Postgres connection
// in order to be used by the other functions of the package.
func openConnection() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func exists(username string) int {
	username = strings.ToLower(username)

	db, err := openConnection()

	if err != nil {
		fmt.Println(err)
		return -1
	}

	defer db.Close()

	UserID := -1
	statement := fmt.Sprintf(`SELECT "id" FROM "users" WHERE username='%s'`, username)
	rows, err := db.Query(statement)

	for rows.Next() {
		var id int
		err = rows.Scan(&id)

		if err != nil {
			fmt.Println("Scan: ", err)
			return -1
		}
		UserID = id
	}
	defer rows.Close()

	return UserID
}

func AddUser(d UserData) int {
	d.Username = strings.ToLower(d.Username)
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := exists(d.Username)

	if userID != -1 {
		fmt.Println("User already exists: ", d.Username)
	}

	insertStatement := fmt.Sprintf(`INSERT INTO "users" ("username") values ($1)`)
	_, err = db.Exec(insertStatement, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	userID = exists(d.Username)
	if userID == -1 {
		fmt.Println("Insert failed")
		return userID
	}

	insertStatement = `INSERT INTO "userdata" ("userid", "name", "surname", "description") values ($1, $2, $3, $4)`
	_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return userID
}
func DeleteUser(id int) error {
	db, err := openConnection()
	if err != nil {
		return err
	}

	defer db.Close()

	statement := fmt.Sprintf(`SELECT "username" FROM "user" WHERE id = %d`, id)

	rows, err := db.Query(statement)

	if err != nil {
		return err
	}
	var username string
	for rows.Next() {
		err = rows.Scan(username)
		if err != nil {
			return err
		}
	}

	defer rows.Close()

	if exists(username) != id {
		return fmt.Errorf("User with ID: %d does not exist", id)
	}

	deleteStatement := fmt.Sprintf(`DELETE FROM "userdata" WHERE "userid"=$1`)

	_, err = db.Exec(deleteStatement, id)

	if err != nil {
		return err
	}
	return nil
}

func ListUser() ([]UserData, error) {
	db, err := openConnection()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	data := []UserData{}
	statement := fmt.Sprintf(`SELECT "id", "username", "name", "surname", "description" FROM "users", "userdata" WHERE users.id=userdata.userid`)

	rows, err := db.Query(statement)

	if err != nil {
		return data, err
	}

	for rows.Next() {
		var id int
		var username string
		var name string
		var surname string
		var description string

		err = rows.Scan(&id, &username, &name, &surname, &description)

		if err != nil {
			return data, err
		}

		temp := UserData{
			ID:          id,
			Username:    username,
			Name:        name,
			Surname:     surname,
			Description: description,
		}

		data = append(data, temp)
	}
	defer rows.Close()

	return data, nil
}

func UpdateUser(d UserData) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	userId := exists(d.Username)

	if userId == -1 {
		return fmt.Errorf("User do not exist")
	}
	d.ID = userId
	updateStatement := `UPDATE "userdata" SET "name"=$1, surname=$2, description=$3 WHERE "userid"=$4`

	_, err = db.Exec(updateStatement, d.Name, d.Surname, d.Description, d.ID)
	if err != nil {
		return err
	}

	return nil
}
