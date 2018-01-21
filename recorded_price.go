package main

// recordedPrice is the market value of BTC on a given day
// example {Date  Low} {Dec 31 2017  12755.60}
type recordedPrice struct {
	Date     string
	LowPrice string
}

func createRecordedPrice(row []string) recordedPrice {
	// get date, lowPrice values from row, create recordedPrice
	record := recordedPrice{Date: row[0], LowPrice: row[3]}
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
