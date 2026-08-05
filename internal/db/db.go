package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func FromDb(){
	db, err := sql.Open("sqlite3", "../../internal/db/test.db")
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	sqlStmt := `CREATE TABLE IF NOT EXISTS users ( 
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT
		);`
	
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'user' created sucessfully")

	fmt.Println("from interna db db.go file")

	
}