package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Read(path string, argOpts ...Option) (map[string][]any, []string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	opts := defaultOptions()
	for _, opt := range argOpts {
		opt(&opts)
	}

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return nil, nil, fmt.Errorf("csv: empty file")
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
		return nil, nil, err
	}

	r := csv.NewReader(f)
	r.Comma = delimiter
	r.TrimLeadingSpace = true

	records, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	if len(records) == 0 {
		return nil, nil, fmt.Errorf("csv: empty file")
	}

	headers := records[0]
	lenHeaders := len(headers)
	if lenHeaders == 0 {
		return nil, nil, fmt.Errorf("csv: no columns found")
	}

	data := make(map[string][]any, lenHeaders)

	for i, row := range records[1:] {
		lenRow := len(row)
		if lenRow != lenHeaders {
			return nil, nil, fmt.Errorf("csv: row %d has %d columns, expected %d", i+1, lenRow, lenHeaders)
		}

		for j, cell := range row {
			col := headers[j]
			data[col] = append(data[col], inferType(cell))
		}
	}

	for _, h := range headers {
		hasFloat := false
		for _, v := range data[h] {
			switch v.(type) {
			case float64:
				hasFloat = true
			}

			if hasFloat {
				break
			}
		}

		if hasFloat {
			for i, v := range data[h] {
				switch n := v.(type) {
				case int64:
					data[h][i] = float64(n)
				}
			}
		}
	}

	return data, headers, nil
}
