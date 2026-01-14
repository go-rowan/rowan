package rowan

import (
	"context"

	"github.com/go-rowan/rowan/internal/sheets"
	"github.com/go-rowan/rowan/table"
)

type SheetsOption = sheets.Option

func FromSheets(ctx context.Context, spreadsheet string, options ...SheetsOption) (*Table, error) {
	data, columns, err := sheets.Read(ctx, spreadsheet, options...)
	if err != nil {
		return nil, err
	}

	return table.New(data, columns)
}
