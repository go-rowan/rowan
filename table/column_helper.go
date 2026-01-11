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

func toFloat64(v any) (float64, bool) {
	switch x := v.(type) {
	case int:
		return float64(x), true
	case int32:
		return float64(x), true
	case int64:
		return float64(x), true
	case float32:
		return float64(x), true
	case float64:
		return x, true
	default:
		return 0, false
	}
}
