package database

import (
	"log"
	"recomendder-go/calc"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Conn *gorm.DB
}

type Document struct {
	ID uint
	Content string
	Title string
}

type TermFreq struct {
	ID uint
	DocumentId uint
	Term string
	Tf float64
}

func NewDb(path string) *DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{Logger: logger.Default})
	if err != nil {
		log.Fatalf("Cannot create sqlite 3 connection %s", err.Error())
	}
	return &DB{
		Conn: db,
	}
}

func (db *DB) Table() {
	db.Conn.AutoMigrate(&Document{}, &TermFreq{})
}

func (db *DB) InsertDocument(content string, title string, termFreq map[string]uint, freqTotal uint) {
	
	// sqlStmt := insertTable(documentTable, "content, title")
	// result, err := db.Conn.Exec(sqlStmt)

	// handleErr(err, sqlStmt)

	// id, err := result.LastInsertId()

	// handleErr(err, "Gagal mengambil id")
	document := &Document{
		Content: content,
		Title: title,
	}
	db.Conn.Create(document)

	insertValue := make([]*TermFreq, 0)

	for term, freq := range termFreq {
		tf := calc.CalcTf(freq, freqTotal)
		insertValue = append(insertValue, &TermFreq{
			Term: term,
			Tf: tf,
			DocumentId: document.ID,
		})
	}

	db.Conn.CreateInBatches(insertValue, len(insertValue))


}


func (db *DB) GetTotalDocument(term string){

}

