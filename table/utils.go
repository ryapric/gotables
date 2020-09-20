package table

import (
	"strconv"
)

func (t *Table) ConvertColStringToFloat64(column string) {
	convertedcol := make([]float64, len(t.Data[column]))
	for rownum := range t.Data[column] {
		val, err := strconv.ParseFloat(t.Data[column][rownum], 64)
		if err != nil {
			panic(err)
		}
		convertedcol = append(convertedcol, val)
	}

	delete(t.Data, column)
	t.Data[column] = convertedcol
}
