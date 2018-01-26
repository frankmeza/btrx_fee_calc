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

func insertIntoMongo(fees []recordedFee, prices []recordedPrice) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	feesColl := session.DB("bittrex").C("fees")
	insertFees(fees, feesColl)

	pricesColl := session.DB("bittrex").C("prices")
	insertPrices(prices, pricesColl)

	defer session.Close()

}

func main() {
	// prices
	histDataReader := createCsvReader(HistDataBTC)
	histData, err := histDataReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// priceHistory []recordedPrice
	priceHistory := processHistData(histData)
	// fmt.Println(priceHistory)

	// fees
	btrxReader := createCsvReader(BTRX)
	btrxData, err := btrxReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// feeHistory []recordedFee
	feeHistory := processBtrxData(btrxData)
	// fmt.Println(feeHistory)

	insertIntoMongo(feeHistory, priceHistory)
}
