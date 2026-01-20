package table

import "fmt"

// New constructs a new Table from a map of column names to slices of values.
// All columns must have the same length; otherwise, an error is returned.
//
// Parameters:
//   - data: map[string][]any, where each key is a column name and the value is a slice of values.
//   - columnsOrder (optional): variadic slice specifying the desired order of columns in the Table.
//     If provided and not nil, the Table will use this order. If not provided, the order will follow the order of iteration over the map (which is non-deterministic in Go).
//     Only the first slice is used if multiple slices are provided.
//
// Returns:
//   - *Table: the constructed Table instance
//   - error: error if the data is empty, if a column in columnsOrder is missing,
//            or if column lengths are inconsistent.
func New(data map[string][]any, columnsOrder ...[]string) (*Table, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("table: data is empty")
	}

	var (
		columns []string
		length  = -1
	)

	if len(columnsOrder) > 0 && columnsOrder[0] != nil {
		columns = columnsOrder[0]

		for _, col := range columns {
			values, ok := data[col]
			if !ok {
				return nil, fmt.Errorf("table: column %s not found in data", col)
			}

			lenVal := len(values)
			if length == -1 {
				length = lenVal
			}

			if lenVal != length {
				return nil, fmt.Errorf("table: column %s has length %d, expected %d", col, lenVal, length)
			}
		}
	} else {
		for col, values := range data {
			lenVal := len(values)
			if length == -1 {
				length = lenVal
			}

			if len(values) != length {
				return nil, fmt.Errorf("table: column %s has length %d, expected %d", col, lenVal, length)
			}

			columns = append(columns, col)
		}
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  length,
	}, nil
}

// NewEmptyTable creates and returns a completely empty Table.
// The table has no columns and no rows.
// This can be used as a placeholder or initial value when no column information is available.
//
// Example:
//   t := NewEmptyTable()
//   fmt.Println(t.Length())   // 0
//   fmt.Println(t.Columns())  // []
func NewEmptyTable() *Table {
	return &Table{
		columns: []string{},
		data:    map[string][]any{},
		length:  0,
	}
}

// EmptyTableFrom creates and returns a new empty Table with the same column headers as the given Table. The returned Table has zero rows, but preserves the column names.
// This is useful for operations like SelectRows([]int{}) or filtering that result in no rows.
//
// Example:
//   t2 := EmptyTableFrom(t)
//   fmt.Println(t2.Length())   // 0
//   fmt.Println(t2.Columns())  // same as t.Columns()
func EmptyTableFrom(t *Table) *Table {
	columns := make([]string, 0, len(t.columns))
	data := make(map[string][]any, len(t.columns))

	for _, c := range t.columns {
		columns = append(columns, c)
		data[c] = []any{}
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  0,
	}
}
