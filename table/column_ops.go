package table

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
