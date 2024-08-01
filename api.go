package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	url := "libsql://monkey-otpvayne.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MjIwMzE4NTUsImlkIjoiNWY4Y2IxMzQtNGFhNi00Njk4LWE3Y2MtNGFiZWIxMjY3NjkxIn0.VsgbBn9UwOF1V3JOOSqtJ55y4esS4CAaEThF8M42flsWsN4sXzdfjUCkYnnYul3oSf0GNn56sNxZiKM8OMeYCQ"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

	queryUsers(db)
	createUser(db, 18, "Flaquetix")
}

type User struct {
	ID   int
	Name string
}

func queryUsers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		users = append(users, user)
		fmt.Println(user.ID, user.Name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}
}

func createUser(db *sql.DB, id int, name string) {
	stmt, err := db.Prepare("INSERT INTO users (id,name) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Nuevo usuario insertado correctamente")
}
