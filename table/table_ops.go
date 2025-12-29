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

func (t *Table) Categorize() *Table {
	columnsCount := len(t.columns)
	data := make(map[string][]any, columnsCount*2)
	columns := make([]string, 0, columnsCount*2)

	for _, c := range t.columns {
		originData := t.data[c]

		d := make([]any, len(originData))
		copy(d, originData)
		data[c] = d

		columns = append(columns, c)

		col, _ := t.Col(c)
		if !col.categorical {
			continue
		}

		ctgMap := make(map[any]int)
		ctgData := make([]any, len(originData))
		index := 0

		for i, v := range originData {
			if _, ok := ctgMap[v]; !ok {
				ctgMap[v] = index
				index++
			}
			ctgData[i] = ctgMap[v]
		}

		headerName := c + "_categorized"
		data[headerName] = ctgData
		columns = append(columns, headerName)
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  t.length,
	}
}
