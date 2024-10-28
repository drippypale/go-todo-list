package csvHandler

import (
	"encoding/csv"
	"log"
	"os"
)

func getOrCreateCsv(csvPath string) *os.File {
	if _, err := os.Stat(csvPath); err != nil {
		if _, err := os.Create(csvPath); err != nil {
			log.Fatalf("Can not create the csv file %v\n%s", csvPath, err)
		}
	}

	f, _ := os.Open(csvPath)

	return f
}

func ReadRecords(csvPath string) [][]string {
	f := getOrCreateCsv(csvPath)
	defer f.Close()

	csvF := csv.NewReader(f)

	data, err := csvF.ReadAll()
	if err != nil {
		panic("CSV file is corrupted.")
	}

	return data
}
