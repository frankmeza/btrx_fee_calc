package main

import (
	"encoding/csv"
	"strings"
)

func bittrexCSVReader() *csv.Reader {
	btrxCsv := strings.NewReader(BTRX)
	r := csv.NewReader(btrxCsv)
	return r
}

func histDataBTCCSVReader() *csv.Reader {
	histDataBtcCsv := strings.NewReader(HistDataBTC)
	r := csv.NewReader(histDataBtcCsv)
	return r
}

func main() {

}
