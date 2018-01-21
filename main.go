package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

// recordedPrice is the market value of BTC on a given day
type recordedPrice struct {
	Date     string
	LowPrice string
}

// recordedFee is the fee paid on a given day
type recordedFee struct {
	Date string
	Fee  string
}

// createCsvReader for use with BTRX, HistDataBTC
func createCsvReader(csvDataString string) *csv.Reader {
	data := strings.NewReader(csvDataString)
	r := csv.NewReader(data)
	return r
}

// recordedPrice functions BEGIN

func createAndSaveRecordedPrice(row []string) recordedPrice {
	// get date, lowPrice values from row, create recordedPrice
	record := recordedPrice{Date: row[0], LowPrice: row[3]}
	return record
}

// process getHistDataLowWithDate :: []string -> []recordedPrice
func processHistData(histData [][]string) []recordedPrice {
	var dataContainer []recordedPrice
	for _, row := range histData {
		record := createAndSaveRecordedPrice(row)
		dataContainer = append(dataContainer, record)
	}
	return dataContainer
}

// recordedPrice functions END

func main() {
	histDataReader := createCsvReader(HistDataBTC)
	histData, err := histDataReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	output := processHistData(histData)
	fmt.Println(output)
}
