package scale

import (
	"fmt"

	"github.com/go-rowan/rowan/internal/numeric"
	"github.com/go-rowan/rowan/table"
)

// ZScaler performs Z-score standardization (also known as standard scaling) on selected numeric columns of a Table.
//
// Z-score standardization transforms each numeric value x using:
//
//	z = (x - mean) / std
//
// where mean and std (standard deviation) are computed from the training data during the Fit phase.
//
// ZScaler is stateful: it stores per-column mean and standard deviation learned during Fit, and applies them consistently during Transform.
//
// Non-numeric values are left unchanged during transformation.
type ZScaler struct {
	mean map[string]float64
	std  map[string]float64
}

// NewZScaler creates and returns a new ZScaler with empty internal state.
//
// The returned scaler must be fitted using Fit before Transform can be called.
// Each call to Fit overwrites any previously stored statistics.
func NewZScaler() *ZScaler {
	return &ZScaler{
		mean: make(map[string]float64),
		std:  make(map[string]float64),
	}
}

// Fit computes the mean and standard deviation for the specified columns from the provided Table.
//
// Only numeric values are considered when computing statistics.
// If a column contains no numeric values, or has zero standard deviation,
// Fit returns an error.
//
// Fit stores the computed statistics internally and overwrites any previously fitted values for the same columns.
func (s *ZScaler) Fit(t *table.Table, columns ...string) error {
	for _, c := range columns {
		col, err := t.Col(c)
		if err != nil {
			return err
		}

		mean, ok := col.Mean()
		if !ok {
			return fmt.Errorf("fit: column %s has no numeric values", c)
		}

		std, ok := col.Std()
		if !ok || std == 0 {
			return fmt.Errorf("fit: column %s has zero standard deviation", c)
		}

		s.mean[c] = mean
		s.std[c] = std
	}

	return nil
}

// Transform applies Z-score standardization to the specified columns using statistics computed during Fit.
//
// Transform returns a new Table with transformed column values.
// The original Table is not modified.
//
// Each numeric value x is transformed as:
//
//	(x - mean) / std
//
// where mean and std are the values learned during Fit.
// Non-numeric values are left unchanged.
//
// Transform returns an error if Fit has not been called for a column or if the stored standard deviation is zero.
func (s *ZScaler) Transform(t *table.Table, columns ...string) (*table.Table, error) {
	result := t.Clone()

	for _, c := range columns {
		col, err := result.Col(c)
		if err != nil {
			return nil, err
		}

		mean := s.mean[c]
		std := s.std[c]

		if std == 0 {
			return nil, fmt.Errorf("transform: cannot standardize column %s with zero std", c)
		}

		mapped := col.Map(func(v any) any {
			f, ok := numeric.ToFloat64(v)
			if !ok {
				return v
			}
			return (f - mean) / std
		})

		if err := result.ReplaceColumn(c, mapped.Values()); err != nil {
			return nil, err
		}
	}

	return result, nil
}
