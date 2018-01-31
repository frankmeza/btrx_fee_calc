package main

import (
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

// recordedFee is the fee paid on a given day
// example {Closed CommissionPaid} {5/21/2017 8:09:18 AM 0.00041584}
type recordedFee struct {
	Date time.Time
	Fee  float64
}

// createRecordedFee :: []string -> recordedFee
func createRecordedFee(row []string) recordedFee {
	date, _ := time.Parse("1/2/2006", row[8])
	f := strings.TrimSpace(row[5])
	fee, _ := strconv.ParseFloat(f, 64)
	record := recordedFee{Date: date, Fee: fee}
	return record
}

// processBtrxData :: [][]string -> []recordedFee
func processBtrxData(btrxData [][]string) []recordedFee {
	var fees []recordedFee
	for _, row := range btrxData {
		record := createRecordedFee(row)
		fees = append(fees, record)
	}
	return fees
}

// insertFees :: []recordedFee, *mgo.Collection -> void
func insertFees(fees []recordedFee, coll *mgo.Collection) {
	for _, f := range fees {
		coll.Insert(&recordedFee{Date: f.Date, Fee: f.Fee})
	}
}
