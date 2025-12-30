package table

import "fmt"

func New(data map[string][]any, columnsOrder ...[]string) (*Table, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("table: data is empty")
	}

	var (
		columns []string
		length  = -1
	)

	if len(columnsOrder) > 0 && columnsOrder[0] != nil {
		columns = columnsOrder[0]

		for _, col := range columns {
			values, ok := data[col]
			if !ok {
				return nil, fmt.Errorf("table: column %s not found in data", col)
			}

			lenVal := len(values)
			if length == -1 {
				length = lenVal
			}

			if lenVal != length {
				return nil, fmt.Errorf("table: column %s has length %d, expected %d", col, lenVal, length)
			}
		}
	} else {
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
	}

	return &Table{
		columns: columns,
		data:    data,
		length:  length,
	}, nil
}

func NewFromStructs[T any](rows []T) (*Table, error) {
	return nil, nil
}
