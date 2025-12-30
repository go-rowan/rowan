package table

import "fmt"

type Column struct {
	name        string
	data        []any
	categorical bool
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
	originData, ok := t.data[name]
	if !ok {
		return nil, fmt.Errorf("column %s not found", name)
	}

	data := make([]any, len(originData))
	copy(data, originData)

	col := &Column{
		name: name,
		data: data,
	}

	col.categorical = inferCategorical(col.data, 3)
	return col, nil
}

func (t *Table) MustCol(name string) *Column {
	col, err := t.Col(name)
	if err != nil {
		return &Column{name: name, data: []any{}}
	}
	return col
}
