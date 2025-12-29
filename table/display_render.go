package table

import (
	"fmt"
	"strings"
)

func displayByIndexes(t *Table, indexes []int) {
	if t == nil {
		fmt.Println("nil")
		return
	}

	if len(indexes) == 0 {
		fmt.Println("-- empty --")
		return
	}

	widths := columnWidths(t, indexes)

	var sb strings.Builder

	separator := renderSeparator(t.Columns(), widths)

	sb.WriteString(separator)
	sb.WriteString("\n")

	sb.WriteString(renderHeader(t.Columns(), widths))
	sb.WriteString("\n")

	sb.WriteString(separator)
	sb.WriteString("\n")

	for _, i := range indexes {
		sb.WriteString(renderRow(t, i, widths))
		sb.WriteString("\n")
	}

	sb.WriteString(separator)
	sb.WriteString("\n")

	fmt.Println(sb.String())
}
