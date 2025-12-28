package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Read(path string) (map[string][]any, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
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
