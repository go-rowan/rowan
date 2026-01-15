package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExcelSource struct {
	path    string
	rangeA1 string
	opts    options
}

func NewExcelSource(path string, argOpts ...Option) (*ExcelSource, error) {
	o := options{
		rangeA1: "Sheet1", // default
	}
	for _, arg := range argOpts {
		arg(&o)
	}

	return &ExcelSource{
		path:    path,
		rangeA1: o.rangeA1,
	}, nil
}

func (s *ExcelSource) Read() ([]string, [][]string, error) {
	// This file uses the Excelize library (github.com/qax-os/excelize), licensed under BSD 3-Clause.
	f, err := excelize.OpenFile(s.path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	sheet := s.rangeA1
	if sheet == "" {
		sheets := f.GetSheetList()
		if len(sheets) == 0 {
			return nil, nil, fmt.Errorf("excel: file has no sheets")
		}
		sheet = sheets[0]
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, nil, err
	}

	if len(rows) == 0 {
		return nil, nil, fmt.Errorf("excel: sheet %s is empty", sheet)
	}

	return rows[0], rows[1:], nil
}
