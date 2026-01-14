package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type CSVSource struct {
	path string
	opts options
}

func NewCSVSource(path string, argOpts ...Option) *CSVSource {
	opts := defaultOptions()
	for _, opt := range argOpts {
		opt(&opts)
	}

	return &CSVSource{
		path: path,
		opts: opts,
	}
}

func (s *CSVSource) Read() ([]string, [][]string, error) {
	f, err := os.Open(s.path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return nil, nil, fmt.Errorf("csv: empty file")
	}
	headerLine := scanner.Text()

	delimiter := s.opts.comma
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
	if len(headers) == 0 {
		return nil, nil, fmt.Errorf("csv: no columns found")
	}

	rows := records[1:]

	return headers, rows, nil
}
