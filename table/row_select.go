package table

import "math/rand"

const defaultDisplayRows = 5

func (t *Table) fetchRows(indexes []int) *Table {
	indexesCount := len(indexes)
	data := make(map[string][]any, len(t.data))

	for col, values := range t.data {
		colData := make([]any, 0, indexesCount)
		for _, index := range indexes {
			colData = append(colData, values[index])
		}
		data[col] = colData
	}

	columns := make([]string, len(t.columns))
	copy(columns, t.columns)

	return &Table{
		columns: columns,
		data:    data,
		length:  indexesCount,
	}
}

func firstIndexes(n, length int) []int {
	if n <= 0 || length <= 0 {
		return nil
	}

	if n > length {
		n = length
	}

	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}

	return indexes
}

func lastIndexes(n, length int) []int {
	if n <= 0 || length <= 0 {
		return nil
	}

	if n > length {
		n = length
	}

	start := length - n

	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = start + i
	}

	return indexes
}

func sampleIndexes(n, length int) []int {
	if n <= 0 || length <= 0 {
		return nil
	}

	if n > length {
		n = length
	}

	perm := rand.Perm(length)
	return perm[:n]
}
