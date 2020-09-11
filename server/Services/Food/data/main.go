package main

import (
	//"bufio"
	"encoding/json"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("RAW_recipes.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	mapCsv := CSVToMap(csvfile)
	json, _ := json.Marshal(mapCsv[0])
	fmt.Printf("data: %v", string(json))
}

// CSVToMap takes a reader and returns an array of dictionaries, using the header row as the keys
func CSVToMap(reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
			continue
		}

		dict := map[string]string{}
		for i := range header {
			dict[header[i]] = record[i]
		}
		rows = append(rows, dict)
	}
	return rows
}