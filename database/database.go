package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

func NewDb(path string) *DB{
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Cannot create sqlite 3 connection %s", err.Error())
	}
	return &DB{
		Conn: db,
	}
}

func (db *DB) Table() {
	sqlstmt := createTable("documents", "id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, content TEXT, title TEXT")

	_, err := db.Conn.Exec(sqlstmt)
	
	handleErr(err, sqlstmt)

	sqlstmt = createTable("termfreq", `id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
										document_id INTEGER NOT NULL, 
										term TEXT, 
										freq INTEGER,
										FOREIGN KEY (document_id) REFERENCES documents(id)
										`)

	_, err = db.Conn.Exec(sqlstmt)

	handleErr(err, sqlstmt)
}

func handleErr(err error, query string){
	if err != nil {
		log.Fatalf("Error: %s, query : %s",err, query )
	}
}

func createTable(tableName string, strucutre string) string {
	return fmt.Sprintf("CREATE TABLE %s (%s)", tableName, strucutre)
}
