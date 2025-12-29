package table

import "fmt"

func (t *Table) Display() {
	if t == nil {
		fmt.Println("nil")
		return
	}

	indexes := firstIndexes(t.Len(), t.Len())
	displayByIndexes(t, indexes)
}
