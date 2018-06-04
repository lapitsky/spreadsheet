package csvparser

import (
	"encoding/csv"
	"fmt"
	"io"
)

// Read function reads CSV file from reader
func Read(file io.Reader) (records [][]string, err error) {
	var csvRecords [][]string
	var csvErr error

	reader := csv.NewReader(file)
	if csvRecords, csvErr = reader.ReadAll(); csvErr != nil {
		err = fmt.Errorf("Could not parse csv file: %s", csvErr)
	}

	return csvRecords, err
}
