package table

import (
	"encoding/csv"
	"os"
)

// Table is the highest-level struct. Tables are treated as maps
type Table struct {
	ColNames []string
	Data map[string][]string
}

// ReadCSV will read a CSV file into a Table struct
func ReadCSV(filepath string) Table {
	csvFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	
	table := tabulateCSV(records)
	return table
}

// tabulateCSV will iterate through what csv.ReadAll() returns, and insert values
// into a Table struct
func tabulateCSV(records [][]string) Table {
	colNames := records[0][:]
	tableRaw := make(map[string][]string)
	// Pad map with colNames
	for _, colName := range colNames {
		tableRaw[colName] = nil
	}


	// Iterate through each record, then each field, appending to the Table map.
	// This loop iterator nomenclature relies on the csv Reader failing if the
	// structure of the read data is bad, and also that ordering of a map
	// iterator is preserved.
	for rowIndex := range records {
		for colIndex, colName := range colNames {
			if rowIndex == 0 {
				continue
			}
			tableRaw[colName] = append(tableRaw[colName], records[rowIndex][colIndex])
		}
	}
	table := Table{ColNames: colNames, Data: tableRaw}
	return table
}
