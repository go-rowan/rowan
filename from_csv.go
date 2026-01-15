package rowan

import (
	"github.com/go-rowan/rowan/internal/csv"
	"github.com/go-rowan/rowan/table"
)

// CSVOption is an alias of csv.Option used to configure CSV reading behavior.
//
// This alias allows CSV-related options to be exposed through the rowan package without requiring users to import the internal csv package directly.
type CSVOption = csv.Option

// FromCSV reads a CSV file and constructs a Table from its contents.
//
// The CSV file is parsed into column-oriented data, where each column is inferred to have a consistent type across all rows.
// The original column order from the CSV header is preserved.
//
// Optional CSVOption values can be provided to customize parsing behavior such as delimiter selection.
func FromCSV(path string, opts ...CSVOption) (*Table, error) {
	data, columns, err := csv.Read(path, opts...)
	if err != nil {
		return nil, err
	}

	return table.New(data, columns)
}
