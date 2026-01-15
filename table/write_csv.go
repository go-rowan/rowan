package table

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// WriteCSV writes the Table to a CSV file specified by filename.
//
// The function creates a temporary file in the same directory as filename and writes all table data to it. If writing succeeds, the temporary file is renamed to the target filename. If any error occurs during writing or flushing, the temporary file is removed to avoid leaving a partial/corrupt file.
//
// If a file with the specified name already exists, it may be overwritten.
//
// The CSV output format:
//   - The first row contains the column headers in the order defined in the Table.
//   - Each subsequent row contains the corresponding values of the Table as strings.
//   - Missing values are written as empty strings.
//
// Error conditions:
//   - if the table has no data or no columns
//   - if creating the temporary file fails
//   - if writing the header or any row fails
//   - if flushing the writer fails
//   - if closing or renaming the temporary file fails
func (t *Table) WriteCSV(filename string) error {
	if t.length == 0 || len(t.columns) == 0 {
		return fmt.Errorf("table: no data to write")
	}

	dir := filepath.Dir(filename)
	tempFile, err := os.CreateTemp(dir, "*.tmp")
	if err != nil {
		return fmt.Errorf("table: failed creating temp file: %w", err)
	}
	tempName := tempFile.Name()
	defer func() {
		tempFile.Close()

		if _, err := os.Stat(tempName); err == nil {
			os.Remove(tempName)
		}
	}()

	writer := csv.NewWriter(tempFile)

	// header
	if err := writer.Write(t.Columns()); err != nil {
		return fmt.Errorf("table: failed writing header: %w", err)
	}

	// rows
	for i := 0; i < t.length; i++ {
		row := make([]string, len(t.columns))

		for j, column := range t.columns {
			values := t.data[column]
			if i < len(values) {
				row[j] = fmt.Sprint(values[i])
			} else {
				row[j] = ""
			}
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("table: failed writing row %d: %w", i, err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("table: failed flushing writer: %w", err)
	}

	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("table: failed closing temp file: %w", err)
	}

	if err := os.Rename(tempName, filename); err != nil {
		return fmt.Errorf("table: failed renaming temp file: %w", err)
	}

	return nil
}
