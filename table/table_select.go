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
