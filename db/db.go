package db

import (
	"encoding/csv"
	"os"

	log "github.com/sirupsen/logrus"
)

var recordsMap map[string]map[string]string

func Init(dbPath string) {
	log.Info("Initializing database...")
	recordsMap = readCsvFile(dbPath)
	log.Info("Initializing Success!")
}

func Find(playerName string) map[string]string {
	return recordsMap[playerName]
}

func readCsvFile(filePath string) map[string]map[string]string {
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
	recordsMap := make(map[string]map[string]string)
	for _, record := range records {
		newRecord := make(map[string]string)
		for i, header := range headers {
			newRecord[header] = record[i]
		}
		recordsMap[newRecord["Name"]] = newRecord
	}
	return recordsMap
}
