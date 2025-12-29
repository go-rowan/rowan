package table

func (t *Table) MapCol(name string, f func(any) any) (*Table, error) {
	oldCol, err := t.Col(name)
	if err != nil {
		return nil, err
	}

	newCol := oldCol.Map(f)
	columns := make([]string, 0, len(t.columns))

	data := make(map[string][]any, len(t.columns))
	for _, c := range t.columns {
		if c == name {
			data[c] = append([]any(nil), newCol.data...)
		} else {
			values := make([]any, len(t.data[c]))
			copy(values, t.data[c])
			data[c] = values
		}

		columns = append(columns, c)
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  t.length,
	}, nil
}
