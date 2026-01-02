package table

// First returns a new table containing the first n rows of the original table.
//
// If n is not provided, First uses the default number of rows that is 5. If n exceeds the table length, all rows are returned.
//
// The original table is not modified.
func (t *Table) First(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := firstIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}

// Last returns a new table containing the last n rows of the original table.
//
// If n is not provided, Last uses the default number of rows that is 5. If n exceeds the table length, all rows are returned.
//
// The original table is not modified.
func (t *Table) Last(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := lastIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}

// Sample returns a new table containing n randomly sampled rows from the original table.
//
// If n is not provided, Sample uses the default number of rows that is 5. Sampling is performed without replacement.
// If n exceeds the table length, all rows are returned in randomized order.
//
// The original table is not modified.
func (t *Table) Sample(n ...int) *Table {
	rows := defaultDisplayRows
	if len(n) > 0 {
		rows = n[0]
	}

	indexes := sampleIndexes(rows, t.Len())
	return t.fetchRows(indexes)
}
