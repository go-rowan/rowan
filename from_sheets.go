package rowan

import (
	"context"

	"github.com/go-rowan/rowan/internal/sheets"
	"github.com/go-rowan/rowan/table"
)

// SheetsOption is an alias of sheets.Option, re-exported to avoid leaking the internal sheets package while still allowing users to configure FromSheets behavior.
type SheetsOption = sheets.Option

// FromSheets constructs a Table from a Google Sheets document.
//
// The spreadsheet argument may be a Spreadsheet ID or full URL.
// Additional options can be provided to control how the sheet is read (e.g. range, or sheet name).
//
// Internally, this function delegates reading to the sheets package and converts the result into a Table.
func FromSheets(ctx context.Context, spreadsheet string, options ...SheetsOption) (*Table, error) {
	data, columns, err := sheets.Read(ctx, spreadsheet, options...)
	if err != nil {
		return nil, err
	}

	return table.New(data, columns)
}
