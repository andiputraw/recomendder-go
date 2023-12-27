package tf

import (
	"encoding/csv"
	"errors"
	"io"
	"log"

	"recomendder-go/database"
	"recomendder-go/lexer"
	"sync"
)

const TEXT_COLLUMN int = 5
const LINK_COLLUMN int = 3

type TF = map[string]uint

type Doc struct {
	Name  string
	Tf    TF
	Total uint
}

func newDoc(name string, tf TF, total uint) *Doc {
	return &Doc{
		name,
		tf,
		total,
	}
}

func getTf(str string) (TF, uint) {
	tf := make(TF, 0)
	lex := lexer.NewLexer(str)
	term, end := lex.Get()
	for !end {
		count, ok := tf[term]
		if !ok {
			tf[term] = 1
			continue
		}
		tf[term] = count + 1
		term, end = lex.Get()
	}
	return tf, uint(lex.Length)
}

func DocFromCsvMultithread(reader *csv.Reader, workerNum int) []*Doc {
	var wg sync.WaitGroup

	jobs := make(chan []string, workerNum)
	result := make(chan *Doc, workerNum)

	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go processCsv(jobs, result, &wg)
	}

	go func() {
		for {
			record, err := reader.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					close(jobs)
					return
				}
				log.Printf("Error getting data from reader: %s\n", err.Error())
			}
			jobs <- record
		}
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	docs := make([]*Doc, 0)

	for result := range result {
		docs = append(docs, result)
	}

	return docs
}

func processCsv(jobs <-chan []string, result chan<- *Doc, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range jobs {
		name := record[LINK_COLLUMN]
		text := record[TEXT_COLLUMN]
		doc := DocFromString(name, text)
		log.Printf("Completed processing %s\n", name)
		result <- doc
	}

}

func DocFromCsv(reader *csv.Reader) []*Doc {
	_, err := reader.Read()
	if err != nil {
		log.Fatalf("Cannot read csv Headers: %s", err.Error())
	}

	record, err := reader.Read()

	if err != nil {
		log.Fatalf("Cannot read csv: %s", err.Error())
	}
	docs := make([]*Doc, 0)

	for record != nil {
		if len(record) < TEXT_COLLUMN {
			log.Fatalf("Collumn length is less that TEXT_COLLUMN")
		}
		text := record[TEXT_COLLUMN]
		link := record[LINK_COLLUMN]
		tf, length := getTf(text)
		doc := newDoc(link, tf, uint(length))
		docs = append(docs, doc)
		record, err = reader.Read()
		if err != nil {
			break
		}
		log.Printf("Completed indexing %s..\n", record[LINK_COLLUMN])
		
	}
	return docs
}

func DocFromString(name, str string) *Doc {
	tf, total := getTf(str)
	return newDoc(name, tf, total)
}

// It may panic
func DocFromCsvToDb(reader *csv.Reader, db *database.DB){
	_, err := reader.Read()
	if err != nil {
		log.Fatalf("Cannot read csv Headers: %s", err.Error())
	}

	record, err := reader.Read()

	if err != nil {
		log.Fatalf("Cannot read csv: %s", err.Error())
	}

	for record != nil {
		if len(record) < TEXT_COLLUMN {
			log.Fatalf("Collumn length is less that TEXT_COLLUMN")
		}
		text := record[TEXT_COLLUMN]
		link := record[LINK_COLLUMN]
		tf, length := getTf(text)
		doc := newDoc(link, tf, length)
		db.InsertDocument(text, doc.Name, doc.Tf, doc.Total)
		record, err = reader.Read()
		if err != nil {
			break
		}
		log.Printf("Completed indexing %s..\n", record[LINK_COLLUMN])
			
	}
	
}
