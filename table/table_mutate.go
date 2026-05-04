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

	delete(t.data, oldName)
	t.data[newName] = values

	for i, c := range t.columns {
		if c == oldName {
			t.columns[i] = newName
			break
		}
	}

	return nil
}
