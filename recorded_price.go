package main

// recordedPrice is the market value of BTC on a given day
type recordedPrice struct {
	Date     string
	LowPrice string
}

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
