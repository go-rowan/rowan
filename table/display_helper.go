package table

import (
	"fmt"
	"strings"
)

func columnWidths(t *Table, indexes []int) map[string]int {
	widths := make(map[string]int)

	for _, col := range t.Columns() {
		widths[col] = len(col)
	}

	for _, i := range indexes {
		for _, col := range t.Columns() {
			val := fmt.Sprint(t.data[col][i])
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

	for _, col := range t.Columns() {
		val := t.data[col][row]

		if isNumeric(val) {
			var strVal string
			switch v := val.(type) {
			case float64, float32:
				strVal = fmt.Sprintf("%.2f", v)
			default:
				strVal = fmt.Sprint(v)
			}
			sb.WriteString(" " + padCenter(strVal, widths[col]) + " |")
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
