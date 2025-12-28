package table

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
