package rowan

import (
	"github.com/go-rowan/rowan/internal/excel"
	"github.com/go-rowan/rowan/table"
)

// ExcelOption is an alias for excel.Option, allowing users to pass configuration options to control how Excel files are read without importing the internal excel package.
type ExcelOption = excel.Option

// FromExcel reads an Excel file from the specified path and constructs a Table from its contents. The first row of the sheet is treated as the header (column names).
//
// Optional ExcelOption values can be provided to customize behavior, such as selecting a specific sheet or A1 range.
//
// FromExcel returns a Table with parsed data, or an error if reading or parsing fails.
func FromExcel(path string, argOpts ...ExcelOption) (*Table, error) {
	data, columns, err := excel.Read(path, argOpts...)
	if err != nil {
		return nil, err
	}

	return table.New(data, columns)
}
