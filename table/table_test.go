package table

import (
	"testing"
)

var table Table = ReadCSV("../testdata/iris.csv")

func TestReadCSV(t *testing.T) {
	var got interface{}
	var want interface{}

	colnames := []string{
		"Sepal.Length",
		"Sepal.Width",
		"Petal.Length",
		"Petal.Width",
		"Species",
	}

	// Right number of columns
	got = len(table.Colnames)
	want = len(colnames)
	if got != want {
		t.Errorf("Not enough column names (got: %d, want: %d", got, want)
	}

	// Right column names
	for colIndex := range colnames {
		got = table.Colnames[colIndex]
		want = colnames[colIndex]
		if got != want {
			t.Errorf("Unexpected column name at position (got: %s, want: %s)", got, want)
		}
	}

	// Right number of rows
	for _, colName := range colnames {
		got = len(table.Data[colName])
		want = 150
		if got != want {
			t.Errorf("Inconsistent row count for column %s (got: %d, want: %d)", colName, got, want)
		}
	}
}

func TestMultiply(t *testing.T) {
	tbl := table
	tbl.Multiply("result", "Sepal.Length", "Sepal.Width")
	got1 := tbl.Data["result"][:5]
	want1 := []int{1, 2, 3, 4, 5}
	if got1 != want1 {
		t.Errorf("Inconsistent Table.Multiply result (got: %v, want: %v", got1, want1)
	}
}
