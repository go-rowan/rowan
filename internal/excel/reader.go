package excel

import "github.com/go-rowan/rowan/internal/parser"

func Read(path string, argOpts ...Option) (map[string][]any, []string, error) {
	source, err := NewExcelSource(path, argOpts...)
	if err != nil {
		return nil, nil, err
	}

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
