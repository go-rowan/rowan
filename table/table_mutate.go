package table

import "fmt"

// RenameColumn renames oldName to newName.
//
// It returns an error if oldName does not exist or newName is already in use.
// It is a no-op if oldName and newName are the same.
func (t *Table) RenameColumn(oldName, newName string) error {
	if newName == oldName {
		return nil
	}

	values, ok := t.data[oldName]
	if !ok {
		return fmt.Errorf("rename: column %s does not exist", oldName)
	}

	if _, exists := t.data[newName]; exists {
		return fmt.Errorf("rename: column %s already exists", newName)
	}

	renameColumn(t, oldName, newName, values)

	return nil
}

func renameColumn(t *Table, oldName, newName string, values []any) {
	delete(t.data, oldName)
	t.data[newName] = values

	for i, c := range t.columns {
		if c == oldName {
			t.columns[i] = newName
			break
		}
	}
}

// RenameColumns renames multiple columns based on the provided map (oldName -> newName).
//
// If any rename is invalid, no changes are applied to the table.
func (t *Table) RenameColumns(nameMap map[string]string) error {
	usedColumns := make(map[string]struct{})
	originalValues := make(map[string][]any)

	for _, c := range t.columns {
		if _, isAssigned := nameMap[c]; !isAssigned {
			usedColumns[c] = struct{}{}
		}
	}

	for oldName, newName := range nameMap {
		if oldName == newName {
			continue
		}

		values, ok := t.data[oldName]
		if !ok {
			return fmt.Errorf("rename: column %s does not exist", oldName)
		}
		originalValues[oldName] = values

		if _, exists := usedColumns[newName]; exists {
			return fmt.Errorf("rename: column name collision for %s", newName)
		}

		usedColumns[newName] = struct{}{}
	}

	for oldName, newName := range nameMap {
		if oldName == newName {
			continue
		}

		renameColumn(t, oldName, newName, originalValues[oldName])
	}

	return nil
}
