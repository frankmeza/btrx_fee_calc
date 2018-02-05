package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type recordedTotal struct {
	NumberOfFees int
	TotalUSD     float64
	TotalBTC     float64
}

// func getUSDFromBTCHistory(feesColl *mgo.Collection, pricesColl *mgo.Collection) recordedTotal {
func getUSDFromBTCHistory(fees []recordedFee, pricesColl *mgo.Collection) recordedTotal {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var index int
	var totalUSD float64
	var totalBTC float64

	for _, fee := range fees {
		var price recordedPrice
		// get the USD/BTC price for the day that the fee was paid
		err = pricesColl.Find(bson.M{"date": fee.Date}).One(&price)
		usdPaid := fee.Fee * price.LowPrice

		index++
		totalBTC += fee.Fee
		totalUSD += usdPaid

		fmt.Printf("On %v\n", fee.Date)
		fmt.Printf("Today I paid %v \n", usdPaid)
		fmt.Printf("fee * price was %v * %v \n", fee.Fee, price.LowPrice)
		fmt.Printf("running total is %v\n\n", totalUSD)
	}
	fmt.Printf("I paid %v fees", index)
	fmt.Printf("I paid %v USD from %v BTC", totalUSD, totalBTC)

	record := recordedTotal{
		NumberOfFees: index,
		TotalUSD:     totalUSD,
		TotalBTC:     totalBTC,
	}
	return record
}
