package table

import "fmt"

// MapCol applies the provided function `f` to all values in the specified column
// of the Table, returning a new Table with the updated column. All other columns
// remain unchanged. The original Table is not modified.
//
// Parameters:
//   - name: the name of the column to transform
//   - f: a function that takes an `any` value and returns a transformed `any` value
//
// Returns:
//   - *Table: a new Table with the transformed column
//   - error: if the specified column does not exist
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

// Categorize returns a new Table where each categorical column produces an additional column with encoded integer values.
//
// For every categorical column, a new column named "<column>_categorized" is appended. Each unique value in the original column is mapped to a zero-based integer, preserving row order. Non-categorical columns are copied as-is.
//
// The original Table is not modified.
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

// Where filters rows of the Table using a predicate function and returns a new Table containing only rows for which the predicate returns true.
//
// The predicate function receives a map representing a single row, where keys are column names and values are the corresponding cell values.
// All columns are preserved in the resulting Table.
//
// Returns an error only if Table construction fails.
func (t *Table) Where(f func(row map[string]any) bool) (*Table, error) {
	filtered := make(map[string][]any)
	cols := t.Columns()

	for _, c := range cols {
		filtered[c] = []any{}
	}

	for i := 0; i < t.Len(); i++ {
		row := make(map[string]any)
		for _, c := range cols {
			row[c] = t.data[c][i]
		}

		if f(row) {
			for _, c := range cols {
				filtered[c] = append(filtered[c], t.data[c][i])
			}
		}
	}

	return New(filtered, cols)
}

func (t *Table) AddColumns(args map[string][]any) (*Table, error) {
	if t == nil {
		return nil, fmt.Errorf("add columns: table is nil")
	}

	argsCount := len(args)
	if argsCount == 0 {
		return nil, fmt.Errorf("add columns: no columns provided")
	}

	for name, values := range args {
		if name == "" {
			return nil, fmt.Errorf("add columns: column name can not be empty")
		}

		if _, exists := t.data[name]; exists {
			return nil, fmt.Errorf("add columns: column %s already exists", name)
		}

		valuesCount := len(values)
		if valuesCount != t.length {
			return nil, fmt.Errorf("add columns: column %s has length %d, expected %d", name, valuesCount, t.length)
		}
	}

	data := make(map[string][]any, len(t.data)+argsCount)
	columns := make([]string, 0, len(t.data)+argsCount)

	for _, c := range t.columns {
		v := make([]any, len(t.data[c]))
		copy(v, t.data[c])

		data[c] = v
		columns = append(columns, c)
	}

	for name, values := range args {
		v := make([]any, len(values))
		copy(v, values)

		data[name] = v
		columns = append(columns, name)
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  t.length,
	}, nil
}

func (t *Table) AddColumn(name string, values []any) (*Table, error) {
	return t.AddColumns(map[string][]any{
		name: values,
	})
}
