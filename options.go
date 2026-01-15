package rowan

import (
	"github.com/go-rowan/rowan/internal/csv"
	"github.com/go-rowan/rowan/internal/excel"
	"github.com/go-rowan/rowan/internal/sheets"
)

// WithDelimiter returns a CSVOption that sets the delimiter rune used when parsing a CSV file.
//
// This is a convenience wrapper around csv.WithDelimiter.
func WithDelimiter(r rune) CSVOption {
	return csv.WithDelimiter(r)
}

// WithSheetsURL configures FromSheets to treat the spreadsheet argument as a full Google Sheets URL instead of a raw spreadsheet ID.
func WithSheetsURL() SheetsOption {
	return sheets.WithSheetsURL()
}

// WithRange specifies the A1 notation range to read from the sheet (e.g. "Sheet1!A1:D100").
func WithSheetsRange(rangeA1 string) SheetsOption {
	return sheets.WithRange(rangeA1)
}

// WithRange specifies the A1 notation range to read from the sheet (e.g. "Sheet1!A1:D100").
func WithExcelRange(rangeA1 string) ExcelOption {
	return excel.WithRange(rangeA1)
}
