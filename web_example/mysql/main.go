package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main(
	// configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	if err := db.Ping();err != nil{
		log.Fatal(err)
	}
	{
		// create a new table
		query := `
			CREATE TABLE users(
				id INT AUTO_INCREMENT ,
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY (id)
			);`
		if _,err := db.Exec(query);err != nil {
			log.Fatal(err)
		}
	}
	{
		// Insert a new user
		username := "johndoe"
		password := "secret"
		createAt := time.Now()

		result , err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil{
			log.Fatal(err)
		}

		id , err := result.LastInsertId()
		fmt.Println(id)
	}
	{
		// query a single user
		var (
			id int
			username string
			password string
			createAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
            log.Fatal(err)
        }
		fmt.Println(id , username,password,createAt)
	}
	{
		// query all users
		type user struct {
			id int
			username string
			password string
			createAt time.Time
		}

		rows , err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil{
			log.Fatal(err)
		}

		defer rows.Close()

		var users []user
		for rows.Next(){
			var u user
			err := rows.Scan(&u.id , &u.username,&u.password,&u.createAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users,u)
		}
		if err := rows.Err(); err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%#v",users)
	}
	{
		_,err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil{
			log.Fatal(err)
		}
	}
)