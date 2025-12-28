package table

type Table struct {
	columns []string
	data    map[string][]any
	length  int
}

func (t *Table) Columns() []string {
	cols := make([]string, len(t.columns))
	copy(cols, t.columns)
	return cols
}

func (t *Table) Len() int {
	return t.length
}
