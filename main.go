package main

import (
	"encoding/csv"
	"strings"
)

type recordedPrice struct {
}

func bittrexCSVReader() *csv.Reader {
	btrxCsv := strings.NewReader(BTRX)
	r := csv.NewReader(btrxCsv)
	return r
}

func histDataBtcCSVReader() *csv.Reader {
	histDataBtcCsv := strings.NewReader(HistDataBTC)
	r := csv.NewReader(histDataBtcCsv)
	return r
}

// process getHistDataLowWithDate :: []string -> RecordedPrice{}
func getHistDataLowWithDate(histData []string) recordedPrice {
}

func main() {

}
