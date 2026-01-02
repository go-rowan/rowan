package rowan

import "github.com/go-rowan/rowan/internal/csv"

// WithDelimiter returns a CSVOption that sets the delimiter rune used when parsing a CSV file.
//
// This is a convenience wrapper around csv.WithDelimiter.
func WithDelimiter(r rune) CSVOption {
	return csv.WithDelimiter(r)
}
