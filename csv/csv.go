package csv

import (
	"bytes"
	"encoding/csv"
)

// ReadCsv is delegated to read into a CSV the content of the bytes in input
// []string -> Headers of the CSV
// [][]string -> Content of the CSV
func ReadCsv(buf []byte, separator rune) ([]string, [][]string) {
	csvReader := csv.NewReader(bytes.NewReader(buf))
	csvReader.Comma = separator
	csvReader.LazyQuotes = true
	csvReader.TrimLeadingSpace = true
	csvData, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	headers := csvData[0]
	// Remove the headers from the row data
	csvData = csvData[1:]
	// Remove the latest element due to headers shift
	return headers, csvData
}
func WriteCsv(headers []string, records [][]string, separator rune) []byte {
	var buff bytes.Buffer
	writer := csv.NewWriter(&buff)
	writer.Comma = separator
	if err := writer.Write(headers); err != nil {
		panic(err)
	}
	if err := writer.WriteAll(records); err != nil {
		panic(err)
	}

	return bytes.TrimSuffix(buff.Bytes(), []byte("\n"))
}
