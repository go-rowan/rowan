package scale

import "github.com/go-rowan/rowan/table"

type Scaler interface {
	Fit(*table.Table, ...string) error
	Transform(*table.Table, ...string) (*table.Table, error)
	Features() []string
	IsFitted() bool
}
