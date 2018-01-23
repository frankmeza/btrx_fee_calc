package main

import "time"

// recordedPrice is the market value of BTC on a given day
// example {Date  Low} {Dec 31 2017  12755.60}
type recordedPrice struct {
	Date     time.Time
	LowPrice string
}

func createRecordedPrice(row []string) recordedPrice {
	date, _ := time.Parse("1/2/2006", row[0])
	record := recordedPrice{Date: date, LowPrice: row[3]}
	return record
}

// process getHistDataLowWithDate :: []string -> []recordedPrice
func processHistData(histData [][]string) []recordedPrice {
	var prices []recordedPrice
	for _, row := range histData {
		record := createRecordedPrice(row)
		prices = append(prices, record)
	}
	return prices
}
