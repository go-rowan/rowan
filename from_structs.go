package rowan

import (
	"fmt"
	"reflect"

	"github.com/go-rowan/rowan/table"
)

// FromStructs constructs a Table from a slice of structs.
//
// Each exported struct field becomes a column in the resulting table.
// The column name is derived from the struct field name by default, or from the `rowan` struct tag if present.
//
// Fields tagged with `rowan:"-"` or unexported fields are ignored.
//
// Example:
//
//	type User struct {
//	    ID    int    `rowan:"id"`
//	    Name  string
//	    Email string `rowan:"-"`
//	}
//
//	tbl, err := rowan.FromStructs([]User{...})
func FromStructs[T any](rows []T) (*Table, error) {
	if len(rows) == 0 {
		return nil, fmt.Errorf("rowan: empty slice")
	}

	t := reflect.TypeOf(rows[0])
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("rowan: expected struct type")
	}

	data := make(map[string][]any)
	columns := []string{}

	// build columns
	for i := 0; i < t.NumField(); i++ {
		columnName, processField := processField(t.Field(i))
		if !processField {
			continue
		}

		columns = append(columns, columnName)
		data[columnName] = []any{}
	}

	// fill rows
	for _, row := range rows {
		v := reflect.ValueOf(row)

		for i := 0; i < t.NumField(); i++ {
			columnName, processField := processField(t.Field(i))
			if !processField {
				continue
			}

			data[columnName] = append(data[columnName], v.Field(i).Interface())
		}
	}

	return table.New(data, columns)
}

func processField(f reflect.StructField) (string, bool) {
	if !f.IsExported() {
		return "", false
	}

	tag := f.Tag.Get("rowan")
	if tag == "-" {
		return "", false
	}

	columnName := f.Name
	if tag != "" {
		columnName = tag
	}

	return columnName, true
}
