package table

import (
	"fmt"

	"github.com/go-rowan/rowan/internal/numeric"
)

// NumericSlice returns a numeric column as []float64 by column index.
func (t *Table) NumericSlice(columnIndex int) ([]float64, error) {
	if columnIndex < 0 || columnIndex >= len(t.columns) {
		return nil, fmt.Errorf("numeric slice: index out of range")
	}

	colName := t.columns[columnIndex]
	values, ok := t.data[colName]
	if !ok {
		return nil, fmt.Errorf("numeric slice: column %s not found in data", colName)
	}

	valuesCount := len(values)
	if valuesCount != t.length {
		return nil, fmt.Errorf("numeric slice: column %s has length of %d, expected %d", colName, valuesCount, t.length)
	}

	result := make([]float64, t.length)

	for i, v := range values {
		f, ok := numeric.ToFloat64(v)
		if !ok {
			return nil, fmt.Errorf("numeric slice: column %s contains non-numeric value at row %d", colName, i)
		}

		result[i] = f
	}

	return result, nil
}

// NumericMatrix returns all columns as [][]float64 (row-major).
func (t *Table) NumericMatrix() ([][]float64, error) {
	columnsCount := len(t.columns)
	if columnsCount == 0 {
		return nil, fmt.Errorf("numeric matrix: table has no columns")
	}

	X := make([][]float64, t.length)
	for i := 0; i < t.length; i++ {
		X[i] = make([]float64, columnsCount)
	}

	for j := 0; j < columnsCount; j++ {
		col, err := t.NumericSlice(j)
		if err != nil {
			return nil, err
		}

		for i := 0; i < t.length; i++ {
			X[i][j] = col[i]
		}
	}

	return X, nil
}
