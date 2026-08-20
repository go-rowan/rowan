---
title: "RenameColumnsFunc()"
---

# RenameColumnsFunc()

## Description

`RenameColumnsFunc()` renames all columns in the table by applying a transformation function to each existing column name.

This operation is strictly **atomic** and delegates its core execution to [`RenameColumns()`](../rename-columns). If the transformation function results in any invalid state or naming collision, the entire operation is aborted, and the table remains completely unchanged.

It is highly useful for systematic schema modifications, such as converting all column names to lowercase, adding prefixes, or applying specific casing conventions.

---

## Signature

```go
func (t *Table) RenameColumnsFunc(fn func(oldName string) string) error
```

---

## Parameters

- ### fn
  `func(oldName string) string`  
  The transformation function that takes the current column name as an argument and returns the new column name.

---

## Return Value

- ### `error`  
  `RenameColumnsFunc()` returns an error if:

  - The transformation function results in duplicate column names (collision)
  - The transformation results in an invalid table state

  If all columns are successfully transformed and renamed, it returns `nil`.

---

## Behavior

- Modifies the table **in-place** only after verifying that the entire transformation is safe
- Guarantees **atomicity**: if a single name collision occurs, no columns are modified
- Automatically maps all existing columns through the provided function before executing the batch rename
- Inherits all safety guarantees from `RenameColumns()`, making it immune to Go's random map iteration issues
- Preserves all underlying row data and the original order of the columns

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

if err := tbl.RenameColumnsFunc(func(oldName string) string {
  return oldName + "_modified"
}); err != nil {
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

[Name_modified Age_modified]
--------------------------------
| Name_modified | Age_modified |
--------------------------------
| Alice         |      30      |
| Bob           |      25      |
--------------------------------
```

---

## Related Methods

- [`(*Table).Columns()`](../columns) - gives a slice of the column name
- [`(*Table).RenameColumn(oldName, newName string)`](../rename-column) - changes the name of a specified column to a new one
- [`(*Table).RenameColumns(nameMap map[string]string)`](../rename-columns) - changes multiple columns name
- [`(*Table).Col()`](methods/col) — extracts a `Column` instance from a `Table`
