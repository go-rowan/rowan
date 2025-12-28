package table

func (t *Table) copy() *Table {
	columnsCount := len(t.columns)
	data := make(map[string][]any, columnsCount)
	columns := make([]string, 0, columnsCount)

	for _, c := range t.columns {
		values := make([]any, len(t.data[c]))
		copy(values, t.data[c])
		data[c] = values

		columns = append(columns, c)
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  t.length,
	}
}
