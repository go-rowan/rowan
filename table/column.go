package table

import "fmt"

// Column represents a single column in a table.
//
// A column holds its name, underlying data, and metadata inferred from its values (such as whether it should be treated as categorical).
type Column struct {
	name        string
	data        []any
	categorical bool
}

// Name returns the name of the column.
func (c *Column) Name() string {
	return c.name
}

// Values returns a copy of the column values.
//
// Modifying the returned slice does not affect the original column data.
func (c *Column) Values() []any {
	vals := make([]any, len(c.data))
	copy(vals, c.data)
	return vals
}

// Col returns a column by name.
//
// The returned column contains a copy of the underlying data, so changes to the column do not mutate the original table. An error is returned if the column does not exist.
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

// MustCol returns a column by name without returning an error.
//
// If the column does not exist, an empty column with the given name is returned. This method is intended for fluent or chainable usage where error handling is intentionally omitted.
func (t *Table) MustCol(name string) *Column {
	col, err := t.Col(name)
	if err != nil {
		return &Column{name: name, data: []any{}}
	}
	return col
}
