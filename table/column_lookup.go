package table

import (
	"fmt"
	"slices"
)

// HasColumn reports whether a column with the given name exists.
func (t *Table) HasColumn(columnName string) bool {
	return slices.Contains(t.columns, columnName)
}

// GetColumnIndex returns the index of the column with the given name.
func (t *Table) GetColumnIndex(columnName string) (int, error) {
	for i, c := range t.columns {
		if c == columnName {
			return i, nil
		}
	}

	return -1, fmt.Errorf("column %s not found", columnName)
}

// MustGetColumnIndex returns the column index or panics if not found.
func (t *Table) MustGetColumnIndex(columnName string) int {
	index, err := t.GetColumnIndex(columnName)
	if err != nil {
		panic(err)
	}

	return index
}
