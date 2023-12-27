package model

import (
	"recomendder-go/calc"
	"recomendder-go/lexer"
	"recomendder-go/tf"
	"sort"
)

type IDF = map[string]int

type ModelInterface interface {
	GetDocs() []*tf.Doc
	GetTotalDocs() int
	GetIdf() IDF
}

type Model struct {
	docs []*tf.Doc
	idf  IDF
}

func (m *Model) GetDocs() []*tf.Doc {
	return m.docs
}

func (m *Model) GetTotalDocs() int {
	return len(m.docs)
}

func (m *Model) GetIdf() IDF {
	return m.idf
}

func find(haystack ModelInterface, needle string) []SearchResult {
	result := make([]SearchResult, 0)

	lex := lexer.NewLexer(needle)
	terms := make([]string, 0)
	term, end := lex.Get()
	for !end {
		terms = append(terms, term)
		term, end = lex.Get()
	}
	total_doc := haystack.GetTotalDocs()

	// (term / total term) * total document, freq
	//

	for _, doc := range haystack.GetDocs() {
		tfidf := 0.0
		for _, term := range terms {
			freq, ok := doc.Tf[term]
			if !ok {
				freq = 0
			}
			freqDoc, ok := haystack.GetIdf()[term]
			if !ok {
				freqDoc = 0
			}
			tfidf += calc.CalcTfIdf(freq, uint(freqDoc), doc.Total, uint(total_doc))
		}
		result = append(result, SearchResult{
			Link:  doc.Name,
			TFIDF: tfidf,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TFIDF < result[j].TFIDF
	})

	return result
}
type SearchResult struct {
	Link  string
	TFIDF float64
}

func take[T any](arr []T, count int) []T {
	return arr[0:count]
}