package rowan

import "github.com/go-rowan/rowan/table"

type Table = table.Table

func New(data map[string][]any) (*Table, error) {
	return table.New(data)
}
