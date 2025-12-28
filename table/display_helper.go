package table

import (
	"fmt"
	"strings"
)

func columnWidths(t *Table, rows int) map[string]int {
	widths := make(map[string]int)

	for _, col := range t.Columns {
		widths[col] = len(col)
	}

	for i := 0; i < rows; i++ {
		for _, col := range t.Columns {
			val := fmt.Sprint(t.Data[col][i])
			lenVal := len(val)
			if lenVal > widths[col] {
				widths[col] = lenVal
			}
		}
	}

	return widths
}

func renderHeader(cols []string, widths map[string]int) string {
	var sb strings.Builder
	sb.WriteString("|")

	for _, col := range cols {
		sb.WriteString(" " + padCenter(col, widths[col]) + " |")
	}

	return sb.String()
}

func padRight(s string, colWidth int) string {
	if len(s) >= colWidth {
		return s
	}
	return s + strings.Repeat(" ", colWidth-len(s))
}

func padCenter(s string, colWidth int) string {
	if len(s) >= colWidth {
		return s
	}

	total := colWidth - len(s)
	left := total / 2
	right := total - left

	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

func renderSeparator(cols []string, widths map[string]int) string {
	var sb strings.Builder
	sb.WriteString("-")

	for _, col := range cols {
		sb.WriteString(strings.Repeat("-", widths[col]+3))
	}

	return sb.String()
}

func renderRow(t *Table, row int, widths map[string]int) string {
	var sb strings.Builder
	sb.WriteString("|")

	for _, col := range t.Columns {
		val := t.Data[col][row]

		if isNumeric(val) {
			sb.WriteString(" " + padCenter(fmt.Sprint(val), widths[col]) + " |")
		} else {
			sb.WriteString(" " + padRight(fmt.Sprint(val), widths[col]) + " |")
		}
	}

	return sb.String()
}

func isNumeric(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
		return true
	default:
		return false
	}
}
