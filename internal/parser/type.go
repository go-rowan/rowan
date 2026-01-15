package parser

import (
	"strconv"
	"strings"
)

func inferType(s string) any {
	s = strings.TrimSpace(s)

	if result, err := strconv.ParseBool(s); err == nil {
		if strings.ToLower(s) == "true" || strings.ToLower(s) == "false" {
			return result
		}
	}

	if result, err := strconv.ParseInt(s, 10, 64); err == nil {
		return result
	}

	if result, err := strconv.ParseFloat(s, 64); err == nil {
		return result
	}

	return s
}
