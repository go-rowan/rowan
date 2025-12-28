package table

import "fmt"

func (t *Table) Select(cols ...string) (*Table, error) {
	argsCount := len(cols)
	if argsCount == 0 {
		return nil, fmt.Errorf("select: no columns specified")
	}

	data := make(map[string][]any, argsCount)
	columns := make([]string, 0, argsCount)

	for _, col := range cols {
		v, ok := t.data[col]
		if !ok {
			return nil, fmt.Errorf("select: column %s does not exist", col)
		}

		values := make([]any, len(v))
		copy(values, v)

		data[col] = values
		columns = append(columns, col)
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  t.length,
	}, nil
}

func (t *Table) Drop(cols ...string) (*Table, error) {
	argsCount := len(cols)
	if argsCount == 0 {
		return t.copy(), nil
	}

	dropSet := make(map[string]struct{}, argsCount)
	for _, c := range cols {
		if _, ok := t.data[c]; !ok {
			return nil, fmt.Errorf("drop: column %s does not exist", c)
		}

		dropSet[c] = struct{}{}
	}

	data := make(map[string][]any)
	columns := make([]string, 0, len(t.columns))

	for _, c := range t.columns {
		if _, drop := dropSet[c]; drop {
			continue
		}

		values := make([]any, len(t.data[c]))
		copy(values, t.data[c])

		data[c] = values
		columns = append(columns, c)
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  t.length,
	}, nil
}
