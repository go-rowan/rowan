package table

import "fmt"

type Column struct {
	name string
	data []any
}

func (c *Column) Name() string {
	return c.name
}

func (c *Column) Values() []any {
	vals := make([]any, len(c.data))
	copy(vals, c.data)
	return vals
}

func (t *Table) Col(name string) (*Column, error) {
	data, ok := t.data[name]
	if !ok {
		return nil, fmt.Errorf("column %s not found", name)
	}

	return &Column{
		name: name,
		data: data,
	}, nil
}
