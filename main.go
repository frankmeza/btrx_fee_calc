package main

import (
	"encoding/csv"
	"log"
	"strings"

	"gopkg.in/mgo.v2"
)

// createCsvReader for use with BTRX, HistDataBTC
func createCsvReader(csvDataString string) *csv.Reader {
	data := strings.NewReader(csvDataString)
	r := csv.NewReader(data)
	return r
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	// prices
	histDataReader := createCsvReader(HistDataBTC)
	// read the csv into histData
	histData, err := histDataReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// ingress [][]string -> []recordedPrice
	priceHistory := processHistData(histData)
	// create a collection
	pricesColl := session.DB("bittrex").C("prices")
	// insert priceHistory into mongo collection
	insertPrices(priceHistory, pricesColl)

	// fees
	btrxReader := createCsvReader(BTRX)
	btrxData, err := btrxReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// feeHistory []recordedFee
	feeHistory := processBtrxData(btrxData)
	feesColl := session.DB("bittrex").C("fees")
	insertFees(feeHistory, feesColl)

	recordedTotalColl := session.DB("bittrex").C("totals")
	total := getUSDFromBTCHistory(feeHistory, pricesColl)

	recordedTotalColl.Insert(&recordedTotal{
		NumberOfFees: total.NumberOfFees,
		TotalUSD:     total.TotalUSD,
		TotalBTC:     total.TotalBTC,
	})
}
