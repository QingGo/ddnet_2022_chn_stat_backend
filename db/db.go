package db

import (
	"encoding/csv"
	"os"

	log "github.com/sirupsen/logrus"
)

var recordsMap map[string][]string
var columns []string

func Init(dbPath string) {
	log.Info("Initializing database...")
	recordsMap, columns = readCsvFile(dbPath)
	log.Info("Initializing Success!")
}

func Find(playerName string) map[string]string {
	records := make(map[string]string)
	for i, column := range columns {
		records[column] = recordsMap[playerName][i]
	}
	return records
}

func readCsvFile(filePath string) (recordsMap map[string][]string, columns []string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file %s: %v", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file as CSV for %s: %v", filePath, err)
	}
	headers := records[0]
	recordsMap = make(map[string][]string)
	for _, record := range records {
		recordsMap[record[0]] = record
	}
	return recordsMap, headers
}
