package sheets

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// SheetsSource implements a tabular data source backed by Google Sheets.
type SheetsSource struct {
	service       *sheets.Service
	spreadsheetID string
	rangeA1       string
	opts          options
}

func NewSheetsSource(ctx context.Context, spreadsheet string, argOpts ...Option) (*SheetsSource, error) {
	o := options{
		rangeA1: "Sheet1", // default
	}
	for _, arg := range argOpts {
		arg(&o)
	}

	spreadsheetID := spreadsheet
	if o.isURL {
		id, err := extractSpreadsheetID(spreadsheet)
		if err != nil {
			return nil, err
		}
		spreadsheetID = id
	}

	service, err := sheets.NewService(ctx, option.WithScopes(sheets.SpreadsheetsReadonlyScope))
	if err != nil {
		return nil, err
	}

	return &SheetsSource{
		service:       service,
		spreadsheetID: spreadsheetID,
		rangeA1:       o.rangeA1,
	}, nil
}

func (s *SheetsSource) Read() ([]string, [][]string, error) {
	if s.service == nil {
		return nil, nil, fmt.Errorf("sheets: service is required")
	}

	resp, err := s.service.Spreadsheets.Values.Get(s.spreadsheetID, s.rangeA1).Do()
	if err != nil {
		return nil, nil, wrapError(err)
	}

	if len(resp.Values) == 0 {
		return nil, nil, fmt.Errorf("sheets: empty sheet")
	}

	headers := make([]string, len(resp.Values[0]))
	for i, h := range resp.Values[0] {
		headers[i] = fmt.Sprint(h)
	}

	headersCount := len(headers)
	if headersCount == 0 {
		return nil, nil, fmt.Errorf("sheets: no columns found")
	}

	rows := make([][]string, len(resp.Values)-1)
	for i, record := range resp.Values[1:] {
		nr := len(record)
		if nr != headersCount {
			return nil, nil, fmt.Errorf("sheets: row %d has %d columns, expected %d", i+1, nr, headersCount)
		}

		row := make([]string, headersCount)
		for j, cell := range record {
			row[j] = fmt.Sprint(cell)
		}

		rows[i] = row
	}

	return headers, rows, nil
}

func wrapError(err error) error {
	msg := err.Error()

	switch {
	case strings.Contains(msg, "SERVICE_DISABLED"):
		return &Error{Kind: ErrAPIDisabled, Err: err}

	case strings.Contains(msg, "quota project"):
		return &Error{Kind: ErrQuotaProject, Err: err}

	case strings.Contains(msg, "ACCESS_TOKEN_SCOPE_INSUFFICIENT"),
		strings.Contains(msg, "insufficientPermissions"):
		return &Error{Kind: ErrUnauthorized, Err: err}

	default:
		return &Error{Kind: ErrUnknown, Err: err}
	}
}
