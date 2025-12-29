package table

func (t *Table) First(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := firstIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}

func (t *Table) Last(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := lastIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}

func (t *Table) Sample(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := sampleIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}
