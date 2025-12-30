package table

func inferCategorical(data []any, maxUnique int) bool {
	if len(data) == 0 {
		return false
	}

	switch data[0].(type) {
	case bool:
		return false
	}

	uniques := make(map[any]struct{})
	for _, d := range data {
		uniques[d] = struct{}{}
		if len(uniques) > maxUnique {
			return false
		}
	}

	switch data[0].(type) {
	case string, int, int64:
		return true
	default:
		return false
	}
}

func isNumericColumn(c *Column) bool {
	for _, v := range c.data {
		if _, ok := numeric(v); ok {
			return true
		}
	}
	return false
}
