package table

import (
	"testing"
)

func TestReadCSV(t *testing.T) {
	var got interface{}
	var want interface{}

	colNames := []string{
		"Sepal.Length", 
		"Sepal.Width",
		"Petal.Length",
		"Petal.Width",
		"Species",
	}

	table := ReadCSV("../testdata/iris.csv")

	// Right number of columns
	got = len(table.ColNames)
	want = len(colNames)
	if got != want {
		t.Errorf("Not enough column names (got: %d, want: %d", got, want)
	}

	// Right column names
	for colIndex := range colNames {
		got = table.ColNames[colIndex]
		want = colNames[colIndex]
		if got != want {
			t.Errorf("Unexpected column name at position (got: %s, want: %s)", got, want)
		}
	}

	// Right number of rows
	for _, colName := range colNames {
		got = len(table.Data[colName])
		want = 150
		if got != want {
			t.Errorf("Inconsistent row count for column %s (got: %d, want: %d)", colName, got, want)
		}
	}
}
