package table

// Table represents a simple in-memory table structure.
// It contains the column names, the underlying data per column, and the number of rows.
type Table struct {
	columns []string
	data    map[string][]any
	length  int
}

// Columns returns a copy of the column names in their current order.
// Modifying the returned slice will not affect the Table.
func (t *Table) Columns() []string {
	cols := make([]string, len(t.columns))
	copy(cols, t.columns)
	return cols
}

// Len returns the number of rows in the Table.
func (t *Table) Len() int {
	return t.length
}

// Clone creates a deep copy of the table.
//
// The cloned table has its own copy of column metadata and underlying data slices, so modifications to the returned table do not affect the original table, and vice versa.
//
// Clone preserves:
//   - column order
//   - column names
//   - row count
//
// This method is intended for non-mutating operations (e.g. normalization, scaling, feature transformation) where a transformed table should be produced without altering the original data.
func (t *Table) Clone() *Table {
	if t == nil {
		return nil
	}

	data := make(map[string][]any, len(t.data))

	columns := make([]string, len(t.columns))
	copy(columns, t.columns)

	for _, c := range t.columns {
		values := make([]any, len(t.data[c]))
		copy(values, t.data[c])
		data[c] = values
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  t.length,
	}
}
