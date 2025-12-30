package rowan

import "github.com/go-rowan/rowan/table"

// Table is an alias of table.Table exposed.
// This allows working with Table directly via the rowan package without importing the internal table package.
//
// Table represents a simple in-memory table structure.
// It contains the column names, the underlying data per column, and the number of rows.
type Table = table.Table

// New creates a new Table from the given column-oriented data.
// This function is a convenience wrapper around table.New.
//
// data: map[string][]any, where each key is a column name and the value is a slice of values.
//
// All columns must have the same length, otherwise an error is returned.
//
// Returns:
//   - *Table: the constructed Table instance
//   - error: error if the data is empty, if a column in columnsOrder is missing,
//            or if column lengths are inconsistent.
func New(data map[string][]any) (*Table, error) {
	return table.New(data)
}
