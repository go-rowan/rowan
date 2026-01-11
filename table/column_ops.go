package table

import "fmt"

// Map applies the provided function `f` to each value in the Column and returns a new Column containing the results. The original Column remains unchanged.
//
// Parameters:
//   - f: a function that takes an `any` value and returns a transformed `any` value.
//
// Returns:
//   - *Column: a new Column with the mapped values.
func (c *Column) Map(f func(any) any) *Column {
	values := make([]any, len(c.data))
	for i, v := range c.data {
		values[i] = f(v)
	}

	return &Column{
		name: c.name,
		data: values,
	}
}

func (c *Column) Normalize() (*Column, error) {
	if c == nil || c.Count() == 0 {
		return nil, fmt.Errorf("normalize: column is empty or nil")
	}

	min, ok := c.Min()
	if !ok {
		return nil, fmt.Errorf("normalize: column is not numeric")
	}
	max, _ := c.Max()

	if min == max {
		return nil, fmt.Errorf("min equals max")
	}

	values := c.Values()
	result := make([]any, 0, len(values))

	for _, v := range values {
		if v == nil {
			result = append(result, nil)
			continue
		}

		x, ok := toFloat64(v)
		if !ok {
			return nil, fmt.Errorf("normalize: non-numeric value encountered")
		}

		n := (x - min) / (max - min)
		result = append(result, n)
	}

	return &Column{
		name: c.name,
		data: result,
	}, nil
}
