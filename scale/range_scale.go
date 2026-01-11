package scale

import (
	"fmt"

	"github.com/go-rowan/rowan/internal/numeric"
	"github.com/go-rowan/rowan/table"
)

// RangeScaler implements range-based feature scaling (often known as Min-Max scaling).
//
// It scales numeric values in specified columns to a normalized range, typically [0, 1], using statistics learned during the Fit phase.
//
// The scaler stores per-column minimum and maximum values computed from a Table, and applies the transformation in a non-mutating manner during Transform.
//
// Formula:
//
//	scaled = (x - min) / (max - min)
//
// RangeScaler is stateful: Fit must be called before Transform.
type RangeScaler struct {
	min map[string]float64
	max map[string]float64
}

// NewRangeScaler creates and initializes a new RangeScaler instance.
//
// The returned scaler has empty internal state and must be fitted using Fit before it can be used to transform data.
func NewRangeScaler() *RangeScaler {
	return &RangeScaler{
		min: make(map[string]float64),
		max: make(map[string]float64),
	}
}

// Fit computes and stores the minimum and maximum values for each specified column.
//
// Fit scans the provided Table and extracts numeric values from the given columns.
// The computed statistics are stored internally and later used by Transform.
//
// Parameters:
//   - t: the input Table used to learn scaling parameters
//   - columns: names of columns to be scaled
//
// Returns:
//   - error if a column does not exist or contains no numeric values
//
// Notes:
//   - Fit does not modify the input Table.
//   - Fit must be called before Transform.
//   - Calling Fit multiple times overwrites previously stored statistics for the specified columns.
func (s *RangeScaler) Fit(t *table.Table, columns ...string) error {
	for _, c := range columns {
		col, err := t.Col(c)
		if err != nil {
			return err
		}

		min, ok := col.Min()
		if !ok {
			return fmt.Errorf("fit: column has no numeric values")
		}
		max, _ := col.Max()

		s.min[c] = min
		s.max[c] = max
	}

	return nil
}

// Transform applies range-based scaling to the specified columns of a Table.
//
// Transform returns a new Table with scaled values, leaving the original Table unchanged.
// Only numeric values are transformed; non-numeric values are preserved as-is.
//
// Parameters:
//   - t: the input Table to be transformed
//   - columns: names of columns to scale
//
// Returns:
//   - *table.Table: a new Table containing the transformed data
//   - error if a column does not exist, has zero range, or was not fitted
//
// Notes:
//   - Transform assumes Fit has already been called for the specified columns.
//   - Scaling is performed using the formula:
//     (x - min) / (max - min)
//   - Columns with zero range (min == max) result in an error to avoid division by zero.
func (s *RangeScaler) Transform(t *table.Table, columns ...string) (*table.Table, error) {
	result := t.Clone()

	for _, c := range columns {
		col, err := result.Col(c)
		if err != nil {
			return nil, err
		}

		min := s.min[c]
		max := s.max[c]
		r := max - min

		if r == 0 {
			return nil, fmt.Errorf("transform: cannot scale column with zero range")
		}

		mapped := col.Map(func(v any) any {
			f, ok := numeric.ToFloat64(v)
			if !ok {
				return v
			}
			return (f - min) / r
		})

		if err := result.ReplaceColumn(c, mapped.Values()); err != nil {
			return nil, err
		}
	}

	return result, nil
}
