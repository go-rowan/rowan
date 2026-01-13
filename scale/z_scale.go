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
// It also stores the name of columns being fitted in a slice of string.
//
// Non-numeric values are left unchanged during transformation.
type ZScaler struct {
	features []string
	mean     map[string]float64
	std      map[string]float64
}

// NewZScaler creates and returns a new ZScaler with empty internal state.
//
// The returned scaler must be fitted using Fit before Transform can be called.
// Each call to Fit overwrites any previously stored statistics.
func NewZScaler() *ZScaler {
	return &ZScaler{
		features: make([]string, 0),
		mean:     make(map[string]float64),
		std:      make(map[string]float64),
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
	if t == nil {
		return fmt.Errorf("fit: table is nil")
	}

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

		s.features = append(s.features, c)
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

		mean, okMean := s.mean[feat]
		std, okStd := s.std[feat]
		if !okMean || !okStd {
			return nil, fmt.Errorf("transform: column %s was not fitted", feat)
		}

		if std == 0 {
			return nil, fmt.Errorf("transform: cannot standardize column %s with zero std", feat)
		}

		mapped := col.Map(func(v any) any {
			f, ok := numeric.ToFloat64(v)
			if !ok {
				return v
			}
			return (f - mean) / std
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
func (s *ZScaler) Features() []string {
	features := make([]string, len(s.features))
	copy(features, s.features)
	return features
}

// Mean returns the mean value learned for the specified column during Fit.
// The boolean return value indicates whether the column was present and fitted.
func (s *ZScaler) Mean(column string) (float64, bool) {
	v, ok := s.mean[column]
	return v, ok
}

// Std returns the standard deviation learned for the specified column during Fit.
// The boolean return value indicates whether the column was present and fitted.
func (s *ZScaler) Std(column string) (float64, bool) {
	v, ok := s.std[column]
	return v, ok
}

// IsFitted reports whether the scaler has been fitted with at least one feature.
//
// It returns true if Fit has been successfully called and at least one column's statistics (mean and standard deviation) have been stored.
// Otherwise, it returns false.
func (s *ZScaler) IsFitted() bool {
	return len(s.features) > 0
}
