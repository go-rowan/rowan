---
title: "RenameColumn()"
---

# RenameColumn()

## Description

`RenameColumn()` changes the name of a specific column in the table from its old name to a new name.

This operation updates the table structure **in-place**. It modifies both the internal data mapping and the column order tracking.

If the old name and the new name are identical, the method performs no operation and returns successfully.

---

## Signature

```go
func (t *Table) RenameColumn(oldName, newName string) error
```

---

## Parameters

- ### oldName
  `string`  
  The current name of the column you want to rename.

- ### newName
  `string`  
  The new name to be assigned to the column.

---

## Return Value

- ### `error`  
  `RenameColumn()` returns an error if:

  - The column specified by `oldName` does not exist in the table
  - The `newName` is already being used by another column in the table

  If the rename succeeds (or if `oldName` and `newName` are identical), it returns `nil`.

---

## Behavior

- Modifies the table **in-place** without creating a new table instance
- Updates the internal data mapping while preserving all row values within the column
- Updates the column order tracking to reflect the new name at the exact same position
- Performs a no-op (does nothing) and returns `nil` if `oldName == newName`
- Prevents data loss by throwing an error if `newName` already exists, avoiding accidental overwrites

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

if err := tbl.RenameColumn("Age", "Age(year)"); err != nil {
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

[Name Age(year)]
---------------------
| Name  | Age(year) |
---------------------
| Alice |    30     |
| Bob   |    25     |
---------------------
```

---

## Related Methods

- [`(*Table).Columns()`](../columns) - gives a slice of the column name
- [`(*Table).RenameColumns()`](../rename-columns) - changes multiple columns name
- [`(*Table).RenameColumnsFunc()`](../rename-columns-func) - changes columns name by a function
- `(*Table).Col(name string)`
