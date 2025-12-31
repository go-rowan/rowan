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
// Parameters:
//   - data: map[string][]any, where each key is a column name and the value is a slice of values.
//   - columnsOrder (optional): variadic slice specifying the desired order of columns in the Table.
//     If provided and not nil, the Table will use this order. If not provided, the order will follow the order of iteration over the map (which is non-deterministic in Go).
//     Only the first slice is used if multiple slices are provided.
//
// All columns must have the same length, otherwise an error is returned.
//
// Returns:
//   - *Table: the constructed Table instance
//   - error: error if the data is empty, if a column in columnsOrder is missing,
//            or if column lengths are inconsistent.
func New(data map[string][]any, columnsOrder ...[]string) (*Table, error) {
	return table.New(data, columnsOrder...)
}
