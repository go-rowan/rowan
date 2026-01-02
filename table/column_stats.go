package table

import (
	"math"
	"sort"
)

// Sum returns the sum of all numeric values in the column.
//
// Non-numeric values are ignored. The second return value indicates whether at least one numeric value was found.
func (c *Column) Sum() (float64, bool) {
	var sum float64
	found := false

	for _, v := range c.data {
		n, ok := numeric(v)
		if !ok {
			continue
		}

		sum += n
		found = true
	}

	return sum, found
}

// Mean returns the arithmetic mean of all numeric values in the column.
//
// Non-numeric values are ignored. The second return value is false if the column contains no numeric values.
func (c *Column) Mean() (float64, bool) {
	var sum float64
	count := 0

	for _, v := range c.data {
		n, ok := numeric(v)
		if !ok {
			continue
		}

		sum += n
		count++
	}

	if count == 0 {
		return 0, false
	}

	return sum / float64(count), true
}

// Min returns the minimum numeric value in the column.
//
// Non-numeric values are ignored. The second return value is false if the column contains no numeric values.
func (c *Column) Min() (float64, bool) {
	var min float64
	firstMark := true

	for _, v := range c.data {
		n, ok := numeric(v)
		if !ok {
			continue
		}

		if firstMark || n < min {
			min = n
			firstMark = false
		}
	}

	return min, !firstMark
}

// Max returns the maximum numeric value in the column.
//
// Non-numeric values are ignored. The second return value is false if the column contains no numeric values.
func (c *Column) Max() (float64, bool) {
	var max float64
	firstMark := true

	for _, v := range c.data {
		n, ok := numeric(v)
		if !ok {
			continue
		}

		if firstMark || n > max {
			max = n
			firstMark = false
		}
	}

	return max, !firstMark
}

// Std returns the sample standard deviation of numeric values in the column.
//
// Non-numeric values are ignored. The second return value is false if fewer than two numeric values are present.
func (c *Column) Std() (float64, bool) {
	var (
		sum         float64
		count       int
		numericData = []float64{}
	)

	for _, v := range c.data {
		n, ok := numeric(v)
		if !ok {
			continue
		}

		sum += n
		count++
		numericData = append(numericData, n)
	}

	if count < 2 {
		return 0, false
	}

	mean := sum / float64(count)

	var squaredDiff float64
	for _, n := range numericData {
		diff := n - mean
		squaredDiff += diff * diff
	}

	variance := squaredDiff / float64(count-1)
	return math.Sqrt(variance), true
}

// Count returns the number of non-missing values in the column.
//
// A value is considered missing if it is nil or an empty string.
func (c *Column) Count() int {
	count := 0

	for _, v := range c.data {
		if v == nil {
			continue
		}

		if s, ok := v.(string); ok && s == "" {
			continue
		}

		count++
	}

	return count
}

// Missing returns the number of missing values in the column.
//
// A value is considered missing if it is nil or an empty string.
func (c *Column) Missing() int {
	missing := 0

	for _, v := range c.data {
		if v == nil {
			missing++
			continue
		}

		if s, ok := v.(string); ok && s == "" {
			missing++
		}
	}

	return missing
}

func numericSlice(data []any) []float64 {
	numSlice := make([]float64, 0, len(data))

	for _, v := range data {
		if n, ok := numeric(v); ok {
			numSlice = append(numSlice, n)
		}
	}

	return numSlice
}

// Quantile returns the q-th quantile of the numeric values in the column.
//
// The parameter q must be in the range [0, 1]. Non-numeric values are ignored. Linear interpolation is used between adjacent values.
// The second return value is false if the column contains no numeric values or if q is out of range.
func (c *Column) Quantile(q float64) (float64, bool) {
	if q < 0 || q > 1 {
		return 0, false
	}

	numSlice := numericSlice(c.data)
	n := len(numSlice)
	if n == 0 {
		return 0, false
	}

	sort.Float64s(numSlice)

	index := q * float64(n-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))

	if lower == upper {
		return numSlice[lower], true
	}

	weight := index - float64(lower)
	return numSlice[lower]*(1-weight) + numSlice[upper]*weight, true
}

// Median returns the median (0.5 quantile) of the numeric values in the column.
//
// Non-numeric values are ignored. The second return value is false if the column contains no numeric values.
func (c *Column) Median() (float64, bool) {
	return c.Quantile(0.5)
}

// Q1 returns the first quartile (0.25 quantile) of the numeric values in the column.
//
// Non-numeric values are ignored. The second return value is false if the column contains no numeric values.
func (c *Column) Q1() (float64, bool) {
	return c.Quantile(0.25)
}

// Q3 returns the third quartile (0.75 quantile) of the numeric values in the column.
//
// Non-numeric values are ignored. The second return value is false if the column contains no numeric values.
func (c *Column) Q3() (float64, bool) {
	return c.Quantile(0.75)
}
