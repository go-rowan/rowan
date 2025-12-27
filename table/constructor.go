package table

import "fmt"

func New(data map[string][]any) (*Table, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("table: data is empty")
	}

	var (
		columns []string
		length  = -1
	)

	for col, values := range data {
		lenVal := len(values)
		if length == -1 {
			length = lenVal
		}

		if len(values) != length {
			return nil, fmt.Errorf("table: column %s has length %d, expected %d", col, lenVal, length)
		}

		columns = append(columns, col)
	}

	return &Table{
		Columns: columns,
		Data:    data,
		Length:  length,
	}, nil
}

func NewFromStructs[T any](rows []T) (*Table, error) {
	return nil, nil
}
