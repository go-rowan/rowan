package table

type Table struct {
	Columns []string
	Data    map[string][]any
	Length  int
}
