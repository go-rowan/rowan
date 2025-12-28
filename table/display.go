package table

import (
	"fmt"
	"strings"
)

func (t *Table) Display() {
	if t == nil {
		fmt.Println("nil")
		return
	}

	widths := columnWidths(t, t.Len())

	var sb strings.Builder

	separator := renderSeparator(t.Columns(), widths)

	sb.WriteString(separator)
	sb.WriteString("\n")

	sb.WriteString(renderHeader(t.Columns(), widths))
	sb.WriteString("\n")

	sb.WriteString(separator)
	sb.WriteString("\n")

	for i := 0; i < t.Len(); i++ {
		sb.WriteString(renderRow(t, i, widths))
		sb.WriteString("\n")
	}

	sb.WriteString(separator)
	sb.WriteString("\n")

	fmt.Println(sb.String())
}
