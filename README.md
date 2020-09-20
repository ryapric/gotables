gotables
========

A Go library for working with tabular data in a slightly-higher-level fashion.
Inspired by higher-level languages such as R, and the Pandas library for Python.

Examples
--------

Here's a toy `main.go` you can run from this repo top-level to play with one of
the test datasets (Fisher's popular iris dataset). It shows how to read in CSV
data, the fields available on the `Table` struct, and some method(s) the `Table`
has.

```go
package main

import (
    "fmt"

    "github.com/ryapric/gotables/table"
)

func main() {
    // ReadCSV() reads a CSV file directly into a Table struct
    tbl := table.ReadCSV("./testdata/iris.csv")

    // Tables have ColNames, RowCount, and Data fields, where Data is a
    // colname-keyed map of the Table contents. You can access columns in the
    // data directly by indexing into the Table.Data map
    fmt.Println(tbl.ColNames)
    fmt.Println(tbl.RowCount)
    fmt.Println(tbl.Data["Species"])
    fmt.Println(tbl.Data["Sepal Length"])

    // Table.MultiplyAcross() stores the rowwise product of your args in a new
    // column in the Table.Data field
    tbl.MultiplyAcross("product", []string{"Sepal Length", "Sepal Width"})
    fmt.Println(tbl.Data["product"])
}
```

Notes/FAQ
---------

>Are the columns in my data typed?

Lol, no. Internally, the entire table is represented as a `map[string][]string`
type, as it's the most flexible and does not bump up against the ugly gotchas of
using `interface{}`. For the methods like `Table.MultiplyAcross()`, the `string`
values are converted on-the-fly into separate `float64` values, mathematically
operated on, converted back into strings, and put into `Table.Data` where they
belong.

I'm 100% open to PRs to make column-typing happen, though!

>Who in the world would ever use this?

I dunno, I worked with a lot of relational data in a past life, and also wanted
to learn Go, so this was what I went with. Honestly, if you're going to be doing
any real work, you should instead use SQLite and the Golang stdlib for local
tabular manipulation, or some client/server DB for heavier stuff.
