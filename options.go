package rowan

import "github.com/go-rowan/rowan/internal/csv"

func WithDelimiter(r rune) CSVOption {
	return csv.WithDelimiter(r)
}
