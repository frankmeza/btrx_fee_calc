package main

import (
	"strconv"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// recordedPrice is the market value of BTC on a given day
// example {Date  Low} {Dec 31 2017  12755.60}
type recordedPrice struct {
	Date     time.Time
	LowPrice float64
}

// createRecordedPrice :: []string -> recordedPrice
func createRecordedPrice(row []string) recordedPrice {
	date, _ := time.Parse("1/2/2006", row[0])
	p := strings.TrimSpace(row[3])
	price, _ := strconv.ParseFloat(p, 64)
	record := recordedPrice{Date: date, LowPrice: price}
	return record
}

// processHistData :: [][]string -> []recordedPrice
func processHistData(histData [][]string) []recordedPrice {
	var prices []recordedPrice
	for _, row := range histData {
		record := createRecordedPrice(row)
		prices = append(prices, record)
	}
	return prices
}

// insertPrices :: []recordedPrice, *mgo.Collection -> void
func insertPrices(prices []recordedPrice, coll *mgo.Collection) {
	for _, p := range prices {
		coll.Insert(&recordedPrice{Date: p.Date, LowPrice: p.LowPrice})
	}
}
