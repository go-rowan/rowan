package rowan

import (
	"github.com/go-rowan/rowan/internal/csv"
	"github.com/go-rowan/rowan/table"
)

func FromCSV(path string) (*Table, error) {
	data, err := csv.Read(path)
	if err != nil {
		return nil, err
	}

	return table.New(data)
}
