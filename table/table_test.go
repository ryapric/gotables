package table

import (
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	tbl := ReadCSV("../testdata/basic.csv")

	var got interface{}
	var want interface{}

	colnames := []string{"a", "b", "c"}

	// Right number of columns
	got = len(tbl.ColNames)
	want = len(colnames)
	if got != want {
		t.Errorf("Not enough column names (got: %d, want: %d", got, want)
	}

	// Right column names
	for colIndex := range colnames {
		got = tbl.ColNames[colIndex]
		want = colnames[colIndex]
		if got != want {
			t.Errorf("Unexpected column name at position (got: %s, want: %s)", got, want)
		}
	}

	// Right number of rows (TODO: make sure this is caught at read-time, as well)
	for _, colName := range colnames {
		got = tbl.RowCount
		want = 3
		if got != want {
			t.Errorf("Inconsistent row count for column %s (got: %d, want: %d)", colName, got, want)
		}
	}
}

func TestMultiply(t *testing.T) {
	tbl := ReadCSV("../testdata/basic.csv")

	var got interface{}
	var want interface{}

	// Subset of operands
	tbl.MultiplyAcross("result", []string{"a", "b"})
	got = tbl.Data["result"]
	want = []string{"2", "20", "56"}
	// Slices can't be compared directly, since their underlying arrays might be
	// different; so reflect.DeepEqual will compare their stored values instead
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Inconsistent Table.Multiply result using two operands (got: %v, want: %v", got, want)
	}

	// All operands
	tbl.MultiplyAcross("result", []string{"a", "b", "c"})
	got = tbl.Data["result"]
	want = []string{"6", "120", "504"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Inconsistent Table.Multiply result using all possible operands (got: %v, want: %v", got, want)
	}
}
