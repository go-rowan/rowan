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
// It also stores the name of columns being fitted in a slice of string.
//
// Formula:
//
//	scaled = (x - min) / (max - min)
//
// RangeScaler is stateful: Fit must be called before Transform.
type RangeScaler struct {
	features []string
	min      map[string]float64
	max      map[string]float64
}

// NewRangeScaler creates and initializes a new RangeScaler instance.
//
// The returned scaler has empty internal state and must be fitted using Fit before it can be used to transform data.
func NewRangeScaler() *RangeScaler {
	return &RangeScaler{
		features: make([]string, 0),
		min:      make(map[string]float64),
		max:      make(map[string]float64),
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
	if t == nil {
		return fmt.Errorf("fit: table is nil")
	}

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

		s.features = append(s.features, c)
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
	if t == nil {
		return nil, fmt.Errorf("transform: table is nil")
	}

	features := columns
	if len(features) == 0 {
		features = s.features
	}
	if len(features) == 0 {
		return nil, fmt.Errorf("transform: no columns specified and scaler has no fitted features")
	}

	result := t.Clone()

	for _, feat := range features {
		col, err := result.Col(feat)
		if err != nil {
			return nil, err
		}

		min, okMin := s.min[feat]
		max, okMax := s.max[feat]
		if !okMin || !okMax {
			return nil, fmt.Errorf("transform: column %s was not fitted", feat)
		}

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

		if err := result.ReplaceColumn(feat, mapped.Values()); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Features returns the list of column names that were fitted by the scaler.
//
// The returned slice represents the features learned during the Fit step and is used as the default set of columns when Transform is called without explicitly specifying columns.
//
// A copy of the underlying slice is returned, so modifying the result will not affect the scaler's internal state.
func (s *RangeScaler) Features() []string {
	features := make([]string, len(s.features))
	copy(features, s.features)
	return features
}

// Min returns the minimum value learned for the specified column during Fit.
// The boolean return value indicates whether the column was present and fitted.
func (s *RangeScaler) Min(column string) (float64, bool) {
	v, ok := s.min[column]
	return v, ok
}

// Max returns the maximum value learned for the specified column during Fit.
// The boolean return value indicates whether the column was present and fitted.
func (s *RangeScaler) Max(column string) (float64, bool) {
	v, ok := s.max[column]
	return v, ok
}

// IsFitted reports whether the scaler has been fitted with at least one feature.
//
// It returns true if Fit has been successfully called and at least one column's statistics (min and max) have been stored.
// Otherwise, it returns false.
func (s *RangeScaler) IsFitted() bool {
	return len(s.features) > 0
}

// Reset clears all learned statistics and fitted features from the scaler.
//
// After calling Reset, the scaler returns to its initial state and must be fitted again using Fit before Transform can be called.
func (s *RangeScaler) Reset() {
	s.features = make([]string, 0)
	s.min = make(map[string]float64)
	s.max = make(map[string]float64)
}
