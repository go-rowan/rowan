package table

import "fmt"

func (t *Table) Overview() {
	if t == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println("Table Overview")
	fmt.Printf("Rows: %d\n", t.Len())
	fmt.Println("Columns:")

	meta := make(map[string][]any)

	meta["Name"] = []any{}
	meta["Type"] = []any{}

	for _, c := range t.Columns() {
		col, _ := t.Col(c)

		meta["Name"] = append(meta["Name"], c)
		meta["Type"] = append(meta["Type"], columnType(col.Values()))
	}

	metaTbl, _ := New(meta)
	metaTbl.Display()
}

func columnType(data []any) string {
	for _, v := range data {
		if v == nil {
			continue
		}

		switch v.(type) {
		case int, int64:
			return "int"
		case float32, float64:
			return "float"
		case bool:
			return "bool"
		case string:
			return "string"
		default:
			return "unknown"
		}
	}
	return "unknown"
}
