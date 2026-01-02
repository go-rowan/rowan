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
