package main

// recordedFee is the fee paid on a given day
type recordedFee struct {
	Date string
	Fee  string
}

func createRecordedFee(row []string) recordedFee {
	record := recordedFee{Date: row[8], Fee: row[5]}
	return record
}

func processBtrxData(btrxData [][]string) []recordedFee {
	var dataContainer []recordedFee
	for _, row := range btrxData {
		record := createRecordedFee(row)
		dataContainer = append(dataContainer, record)
	}
	return dataContainer
}
