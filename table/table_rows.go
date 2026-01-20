package table

import "fmt"

// First returns a new table containing the first n rows of the original table.
//
// If n is not provided, First uses the default number of rows that is 5. If n exceeds the table length, all rows are returned.
//
// The original table is not modified.
func (t *Table) First(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := firstIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}

// Last returns a new table containing the last n rows of the original table.
//
// If n is not provided, Last uses the default number of rows that is 5. If n exceeds the table length, all rows are returned.
//
// The original table is not modified.
func (t *Table) Last(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := lastIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}

// Sample returns a new table containing n randomly sampled rows from the original table.
//
// If n is not provided, Sample uses the default number of rows that is 5. Sampling is performed without replacement.
// If n exceeds the table length, all rows are returned in randomized order.
//
// The original table is not modified.
func (t *Table) Sample(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := sampleIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}

// SelectRows returns a new Table containing only the rows specified by the given indices.
// If any index is out of bounds, an error is returned.
// The original Table is not modified.
//
// Example:
//   t2, err := t.SelectRows([]int{0, 2, 5})
//   if err != nil {
//       // handle error
//   }
func (t *Table) SelectRows(indexes []int) (*Table, error) {
	indexCount := len(indexes)
	if indexCount == 0 {
		return EmptyTableFrom(t), nil
	}

	data := make(map[string][]any, len(t.columns))

	for _, c := range t.columns {
		originCol := t.data[c]
		col := make([]any, indexCount)

		for i, index := range indexes {
			if index < 0 || index >= t.length {
				return nil, fmt.Errorf("select rows: index out of range at the order of %d", i)
			}

			col[i] = originCol[index]
		}

		data[c] = col
	}

	return New(data, t.columns)
}

// MustSelectRows returns a new Table containing only the rows specified by the given indices.
// It panics if any index is out of bounds.
// The original Table is not modified.
//
// Example:
//   t2 := t.MustSelectRows([]int{0, 2, 5}) // panics if index invalid
func (t *Table) MustSelectRows(indexes []int) *Table {
	indexCount := len(indexes)
	if indexCount == 0 {
		return EmptyTableFrom(t)
	}

	data := make(map[string][]any, len(t.columns))
	columns := []string{}

	for _, c := range t.columns {
		originCol := t.data[c]
		col := make([]any, indexCount)

		for i, index := range indexes {
			if index < 0 || index >= t.length {
				panic("select rows: index out of range")
			}

			col[i] = originCol[index]
		}

		data[c] = col
		columns = append(columns, c)
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  indexCount,
	}
}
