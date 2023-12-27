package main

import (
	"encoding/csv"
	"log"
	"os"
	"recomendder-go/database"
	"recomendder-go/tf"
)




func main() {
	file, err := os.Open("dataset/gfg_articles.csv")
	if err != nil {
		log.Fatalf("Cannot open file: %s", err.Error())
	}

	db := database.NewDb("./db.sqlite3")
	
	db.Table()
	defer file.Close()

	reader := csv.NewReader(file)
	log.Printf("Start indexing...\n")
	// docs := tf.DocFromCsvMultithread(reader, 100)
	// docs := tf.DocFromCsv(reader)
	tf.DocFromCsvToDb(reader, db)
	// log.Printf("Index completed\n")
	// log.Printf("Processed %d documents\n", len(docs))

	// idf := make(IDF, 0)
	// for _, doc := range docs {
	// 	for term := range doc.Tf {
	// 		count, ok := idf[term]
	// 		if !ok {
	// 			idf[term] = 1
	// 			continue
	// 		}
	// 		idf[term] = count + 1
	// 	}
	// }
	// model := Model{
	// 	docs,
	// 	idf,
	// }

	// result := find(&model, "golang")
	// top5 := take(result, 5)
	// for i, top := range top5 {
	// 	fmt.Printf("Rank %d : %s\n", i+1, top.link)
	// }
}
