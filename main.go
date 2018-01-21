package main

import (
	"encoding/csv"
	"strings"
)

// type recordedPrice struct {
// }

// createCsvReader for use with BTRX, HistDataBTC
func createCsvReader(csvDataString string) *csv.Reader {
	data := strings.NewReader(csvDataString)
	r := csv.NewReader(data)
	return r
}

// process getHistDataLowWithDate :: []string -> RecordedPrice{}
// func getHistDataLowWithDate(histData []string) recordedPrice {
// }

func main() {
	createCsvReader(BTRX)
}
