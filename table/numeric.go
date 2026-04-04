package table

import (
	"errors"
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

// MustNumericSlice returns a slice of float64 values for the column at the specified index.
//
// This method panics if:
//   - columnIndex is out of range
//   - any value in the column cannot be converted to float64
//
// MustNumericSlice assumes that the Table was constructed correctly:
//   - all columns exist in the table's data map
//   - all columns have the same length as the table
//
// Use this method only when you are certain that the column contains numeric values.
// It is intended for internal or performance-sensitive code where error handling via panic is acceptable.
func (t *Table) MustNumericSlice(columnIndex int) []float64 {
	if columnIndex < 0 || columnIndex >= len(t.columns) {
		panic(errors.New("numeric slice: index out of range"))
	}

	colName := t.columns[columnIndex]
	values, _ := t.data[colName]

	result := make([]float64, t.length)

	for i, v := range values {
		f, ok := numeric.ToFloat64(v)
		if !ok {
			panic(fmt.Errorf("numeric slice: column %s contains non-numeric value at row %d", colName, i))
		}

		result[i] = f
	}

	return result
}

// ColumnToIntSlice retrieves all values from the specified column and converts them to a slice of int.
//
// It can handle various numeric types (int, int64, float64, etc.), but will truncate any decimal points.
//
// Returns an error if the table is nil, the column does not exist, or a value cannot be converted.
func (t *Table) ColumnToIntSlice(column string) ([]int, error) {
	if t == nil || t.data == nil {
		return nil, errors.New("table is nil")
	}

	col, err := t.Col(column)
	if err != nil {
		return nil, err
	}

	data := col.Values()

	intData := make([]int, len(data))

	for i, val := range data {
		floatVal, ok := numeric.ToFloat64(val)
		if !ok {
			return nil, fmt.Errorf("row %d in column %s has unsupported type: %T", i, column, val)
		}

		intData[i] = int(floatVal)
	}

	return intData, nil
}

// MustColumnToIntSlice is a convenience wrapper for ColumnToIntSlice that panics if an error occurs.
//
// Use this only when you are certain the column exists and contains valid numeric data.
func (t *Table) MustColumnToIntSlice(column string) []int {
	intData, err := t.ColumnToIntSlice(column)
	if err != nil {
		panic(err)
	}

	return intData
}
