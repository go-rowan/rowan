package sheets

import (
	"fmt"
	"strings"
)

func extractSpreadsheetID(url string) (string, error) {
	const marker = "/spreadsheets/d/"
	i := strings.Index(url, marker)
	if i == -1 {
		return "", fmt.Errorf("sheets: invalid google sheets URL")
	}

	id := url[i+len(marker):]
	if j := strings.Index(id, "/"); j != -1 {
		id = id[:j]
	}

	if id == "" {
		return "", fmt.Errorf("sheets: empty spreadsheet ID")
	}

	return id, nil
}
