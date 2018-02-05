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

func crunchTheNumbers(feesColl *mgo.Collection, pricesColl *mgo.Collection) recordedTotal {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// get all fees from db
	var fees []recordedFee
	err = feesColl.Find(nil).All(&fees)

	var index int
	var totalUSD float64
	var totalBTC float64

	for _, fee := range fees {
		// get fee in btc, date
		var price recordedPrice
		err = pricesColl.Find(bson.M{"date": fee.Date}).One(&price)
		usdPaid := fee.Fee * price.LowPrice

		// increment
		index++
		totalBTC += fee.Fee
		totalUSD += usdPaid

		fmt.Printf("\n")
		fmt.Printf("On %v\n", fee.Date)
		fmt.Printf("fee * price was %v * %v \n", fee.Fee, price.LowPrice)
		fmt.Printf("Today I paid %v \n", usdPaid)
		fmt.Printf("running total is %v", totalUSD)
		fmt.Println("= = = = = =")

	}
	fmt.Println("I paid fees this many times", index)
	fmt.Println("In total, I paid USD", totalUSD)
	fmt.Println("From BTC", totalBTC)

	record := recordedTotal{
		NumberOfFees: index,
		TotalUSD:     totalUSD,
		TotalBTC:     totalBTC,
	}
	return record
}
