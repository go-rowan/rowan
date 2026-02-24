package table

import "fmt"

// Display prints the table to the standard output.
//
// It renders all rows of the table using the current column order.
// If the table is nil, the string "nil" is printed instead.
func (t *Table) Display() {
	if t == nil {
		fmt.Println("nil")
		return
	}

	indexes := firstIndexes(t.Len(), t.Len())
	displayByIndexes(t, indexes)
}

func (t *Table) DisplayTranspose() {
	if t == nil {
		fmt.Println("nil")
		return
	}

	displayTranspose(t)
}
