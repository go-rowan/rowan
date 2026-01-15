package csv

import (
	"github.com/go-rowan/rowan/internal/parser"
)

func Read(path string, argOpts ...Option) (map[string][]any, []string, error) {
	source := NewCSVSource(path, argOpts...)

	columns, rows, err := source.Read()
	if err != nil {
		return nil, nil, err
	}

	data, err := parser.ParseRows(columns, rows)
	if err != nil {
		return nil, nil, err
	}

	return data, columns, nil
}
