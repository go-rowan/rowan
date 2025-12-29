package table

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
