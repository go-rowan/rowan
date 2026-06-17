---
title: "RenameColumns()"
---

# RenameColumns()

## Description

`RenameColumns()` renames multiple columns simultaneously based on a provided mapping of old names to new names.

This operation is strictly **atomic**. If any single rename pair within the map is invalid, the entire operation fails, and the table remains completely unmodified.

This method safely handles complex operations such as **swapping** column names (e.g., A to B, and B to A) or **chaining** renames (e.g., A to B, and B to C) without data loss or random execution order issues.

---

## Signature

```go
func (t *Table) RenameColumns(nameMap map[string]string) error
```

---

## Parameters

- ### nameMap
  `map[string]string`  
  A map where the keys represent the current column names (`oldName`) and the values represent the desired new column names (`newName`).

---

## Return Value

- ### `error`  
  `RenameColumns()` returns an error if:

  - Any specified `oldName` does not exist in the table
  - Any specified `newName` causes a collision with an existing column (unless that existing column is also being renamed in the same operation)

  If the batch rename succeeds, it returns `nil`.

---

## Behavior

- Modifies the table **in-place** only after all validations have passed successfully
- Guarantees **atomicity**: either all columns are renamed, or none are changed if an error occurs
- Safely handles column name **swapping** and **chaining** independently of Go's random map iteration order
- Preserves all row data and structural positions for the renamed columns
- Skips any pair where `oldName == newName` without triggering an error

---

## Example Usage

```go
tbl, _ := rowan.FromStructs([]struct {
  Name string
  Age  int
}{
  {"Alice", 30},
  {"Bob", 25},
})

fmt.Println(tbl.Columns())
tbl.Display()

nm := map[string]string{
  "Name": "Nombre",
  "Age":  "Edad",
}

if err := tbl.RenameColumns(nm); err != nil {
  panic(err)
}

fmt.Println(tbl.Columns())
tbl.Display()
```

Output:

```
[Name Age]
---------------
| Name  | Age |
---------------
| Alice | 30  |
| Bob   | 25  |
---------------

[Nombre Edad]
-----------------
| Nombre | Edad |
-----------------
| Alice  |  30  |
| Bob    |  25  |
-----------------
```

---

## Related Methods

- [`(*Table).Columns()`](../columns) - gives a slice of the column name
- [`(*Table).RenameColumn(oldName, newName string)`](../rename-column) - changes the name of a specified column to a new one
- [`(*Table).RenameColumnsFunc()`](../rename-columns) - changes columns name by a function
- `(*Table).Col(name string)`
