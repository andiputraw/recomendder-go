package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"recomendder-go/database"
	"recomendder-go/lexer"
	"recomendder-go/tf"
	"sort"
)

type IDF = map[string]int

type Model struct {
	Docs []*tf.Doc
	idf IDF
}

func calc_tf(term string, doc *tf.Doc) float64{
	freq, ok := doc.Tf[term]
	if !ok {
		freq = 0
	}
	return float64(freq) / float64(doc.Total)
}

func calc_idf(term string,total_doc int, idf IDF) float64{
	freq , ok := idf[term]
	if !ok {
		freq = 1
	}
	return math.Log10(float64(total_doc) / float64(freq))
}

func calc_tfidf(term string, total_doc int, doc *tf.Doc, idf IDF) float64{
	return calc_tf(term, doc) * calc_idf(term, total_doc, idf)
}

func find(haystack Model,needle string) []SearchResult{
	result := make([]SearchResult, 0)
	
	lex := lexer.NewLexer(needle)
	terms := make([]string, 0)
	term, end := lex.Get()
	for !end {
		terms = append(terms, term)
		term, end = lex.Get()
	}
	total_doc := len(haystack.Docs)

	for _, doc := range haystack.Docs {
		tfidf := 0.0
		for _, term := range terms {
			tfidf += calc_tfidf(term, total_doc, doc, haystack.idf)
		}
		result = append(result, SearchResult{
			link: doc.Name,
			tfidf: tfidf,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].tfidf < result[j].tfidf
	})

	return result
}

type SearchResult struct {
	link string
	tfidf float64
}

func take[T any](arr []T, count int) []T {
	return arr[0:count]
}

func main() {
	file, err := os.Open("dataset/gfg_articles.csv")
	if err != nil {
		log.Fatalf("Cannot open file: %s", err.Error())
	}

	db := database.NewDb("./db.sqlite3")
	defer db.Conn.Close()
	db.Table()
	defer file.Close()
	
	reader := csv.NewReader(file)
	log.Printf("Start indexing...\n")
	// docs := tf.DocFromCsvMultithread(reader, 100)
	docs := tf.DocFromCsv(reader)
	log.Printf("Index completed\n")
	log.Printf("Processed %d documents\n", len(docs))
	
	idf := make(IDF, 0)
	for _, doc := range docs {
		for term := range doc.Tf {
			count, ok := idf[term]
			if !ok {
				idf[term] = 1
				continue
			}
			idf[term] = count + 1 
		}
	}
	model := Model{
		docs,
		idf,
	}

	result := find(model, "golang")
	top5 := take(result, 5)
	for i, top := range top5 {
		fmt.Printf("Rank %d : %s\n", i + 1, top.link)
	}
}

