package scale

import (
	"fmt"

	"github.com/go-rowan/rowan/table"
)

type RangeScaler struct {
	min map[string]float64
	max map[string]float64
}

func NewRangeScaler() *RangeScaler {
	return &RangeScaler{
		min: make(map[string]float64),
		max: make(map[string]float64),
	}
}

func (s *RangeScaler) Fit(t *table.Table, columns ...string) error {
	for _, c := range columns {
		col, err := t.Col(c)
		if err != nil {
			return err
		}

		min, ok := col.Min()
		if !ok {
			return fmt.Errorf("fit: column has no numeric values")
		}
		max, _ := col.Max()

		s.min[c] = min
		s.max[c] = max
	}

	return nil
}

func (s *RangeScaler) Transform(t *table.Table, columns ...string) (*table.Table, error) {
	return nil, nil
}
