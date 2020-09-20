package table

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Table is the highest-level struct. Tables are treated as maps
type Table struct {
	ColNames []string
	Data     map[string][]string
	RowCount int
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

	// Assign any empty struct fields now that the Table is ready
	table.RowCount = len(table.Data[table.ColNames[0]])

	return table
}

// tabulateCSV will iterate through what csv.ReadAll() returns ([][]string), and
// insert values into a Table struct
func tabulateCSV(records [][]string) Table {
	colnames := records[0][:]
	tableRaw := make(map[string][]string, len(records))

	// Pad map with colnames
	for _, colname := range colnames {
		tableRaw[colname] = nil
	}

	// Iterate through each record, then each field, appending to the Table map.
	// This loop iterator nomenclature relies on the csv Reader failing if the
	// structure of the read data is bad, and also that ordering of a map iterator
	// is preserved.
	for rowIndex := range records {
		for colIndex, colName := range colnames {
			if rowIndex == 0 {
				continue
			}
			tableRaw[colName] = append(tableRaw[colName], records[rowIndex][colIndex])
		}
	}
	table := Table{ColNames: colnames, Data: tableRaw}
	return table
}

// MultiplyAcross multiples columns together, and stores the result in a column
// named `result` (i.e. Table.Data["result"])
func (t *Table) MultiplyAcross(resultCol string, operands []string) {
	res := make([]string, t.RowCount)

	for operandIdx, operand := range operands {
		// Stop when you hit the second-to-last operand, otherwise you'll run out of
		// bounds
		if operandIdx == len(operands)-1 {
			break
		}

		for colname, colvalues := range t.Data {
			if colname != operand {
				continue
			}

			for rownum := range colvalues {
				val1, err := strconv.ParseFloat(t.Data[operand][rownum], 32)
				if err != nil {
					panic(err)
				}
				val2, err := strconv.ParseFloat(t.Data[operands[operandIdx+1]][rownum], 32)
				if err != nil {
					panic(err)
				}

				// First pass should just store the product of the two operands. Any
				// subsequent passes need to multiply your second operand value by the
				// existing value in the Table.
				if operandIdx == 0 {
					resval := strconv.FormatFloat(val1*val2, 'f', -1, 32)
					res[rownum] = resval
				} else {
					oldval, err := strconv.ParseFloat(res[rownum], 32)
					if err != nil {
						panic(err)
					}
					res[rownum] = strconv.FormatFloat(oldval*val2, 'f', -1, 32)
				}
			}
		}
	}

	t.Data[resultCol] = res
}
