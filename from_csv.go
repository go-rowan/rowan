package rowan

import (
	"github.com/go-rowan/rowan/internal/csv"
	"github.com/go-rowan/rowan/table"
)

type CSVOption = csv.Option

func FromCSV(path string, opts ...CSVOption) (*Table, error) {
	data, err := csv.Read(path, opts...)
	if err != nil {
		return nil, err
	}

	return table.New(data)
}
