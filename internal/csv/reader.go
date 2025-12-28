package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Read(path string, argOpts ...Option) (map[string][]any, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	opts := defaultOptions()
	for _, opt := range argOpts {
		opt(&opts)
	}

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return nil, fmt.Errorf("csv: empty file")
	}
	headerLine := scanner.Text()

	delimiter := opts.comma
	if delimiter == 0 {
		commaCount := strings.Count(headerLine, ",")
		semicolonCount := strings.Count(headerLine, ";")

		delimiter = ','
		if semicolonCount > commaCount {
			delimiter = ';'
		}
	}

	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	r.Comma = delimiter
	r.TrimLeadingSpace = true

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("csv: empty file")
	}

	headers := records[0]
	lenHeaders := len(headers)
	if lenHeaders == 0 {
		return nil, fmt.Errorf("csv: no columns found")
	}

	data := make(map[string][]any, lenHeaders)

	for i, row := range records[1:] {
		lenRow := len(row)
		if lenRow != lenHeaders {
			return nil, fmt.Errorf("csv: row %d has %d columns, expected %d", i+1, lenRow, lenHeaders)
		}

		for j, cell := range row {
			col := headers[j]
			data[col] = append(data[col], cell)
		}
	}

	return data, nil
}
