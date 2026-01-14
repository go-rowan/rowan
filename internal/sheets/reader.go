package sheets

import (
	"context"
	"fmt"

	"github.com/go-rowan/rowan/internal/parser"
)

func Read(ctx context.Context, spreadsheet string, argOpts ...Option) (map[string][]any, []string, error) {
	source, err := NewSheetsSource(ctx, spreadsheet, argOpts...)
	if err != nil {
		return nil, nil, err
	}

	columns, rows, err := source.Read()
	if err != nil {
		return nil, nil, err
	}
	columnsCount := len(columns)

	fmt.Printf("columns: %v\n", columns)
	fmt.Printf("rows: %v\n", rows)

	data := make(map[string][]any, columnsCount)

	for i, row := range rows {
		rowsCount := len(row)
		if rowsCount != columnsCount {
			return nil, nil, fmt.Errorf("sheets: row %d has %d columns, expected %d", i+1, rowsCount, columnsCount)
		}

		for j, cell := range row {
			col := columns[j]
			data[col] = append(data[col], parser.InferType(cell))
		}
	}

	for _, c := range columns {
		hasFloat := false
		for _, v := range data[c] {
			if _, ok := v.(float64); ok {
				hasFloat = true
				break
			}
		}

		if hasFloat {
			for i, v := range data[c] {
				if n, ok := v.(int64); ok {
					data[c][i] = float64(n)
				}
			}
		}
	}

	return data, columns, nil
}
