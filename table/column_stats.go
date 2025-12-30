package table

import "math"

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
