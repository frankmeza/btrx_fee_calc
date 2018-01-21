package main

// recordedFee is the fee paid on a given day
// example {Closed CommissionPaid} {5/21/2017 8:09:18 AM 0.00041584}
type recordedFee struct {
	Date string
	Fee  string
}

func createRecordedFee(row []string) recordedFee {
	record := recordedFee{Date: row[8], Fee: row[5]}
	return record
}

func processBtrxData(btrxData [][]string) []recordedFee {
	var fees []recordedFee
	for _, row := range btrxData {
		record := createRecordedFee(row)
		fees = append(fees, record)
	}
	return fees
}
