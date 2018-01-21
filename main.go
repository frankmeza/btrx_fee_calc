package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

// createCsvReader for use with BTRX, HistDataBTC
func createCsvReader(csvDataString string) *csv.Reader {
	data := strings.NewReader(csvDataString)
	r := csv.NewReader(data)
	return r
}

func main() {
	// prices
	histDataReader := createCsvReader(HistDataBTC)
	histData, err := histDataReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	priceHistory := processHistData(histData)
	fmt.Println(priceHistory)

	// fees
	btrxReader := createCsvReader(BTRX)
	btrxData, err := btrxReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	feeHistory := processBtrxData(btrxData)
	fmt.Println(feeHistory)
}
