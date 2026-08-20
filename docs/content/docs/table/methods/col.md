---
title: "Col()"
---

# Col()

## Description

`Col()` retrieves a single column from the table by its name and returns it as a standalone `*Column` instance.

To preserve data integrity, `Col()` creates a shallow copy of the column's underlying slice, ensuring that subsequent mutations or analytical operations on the returned `Column` do not affect the original `Table`.

---

## Signature

```go
func (t *Table) Col(name string) (*Column, error)
```

---

## Parameters

- `name`  
  The exact string identifier of the column to retrieve.

---

## Return Values

- `*Column`  
  A pointer to the newly constructed `Column` instance containing a copy of the target column's data and its inferred categorical status.

- `error`  
  An error is returned if the specified column name does not exist in the table.

---

## Example

```go
tbl, err := rowan.FromExcel("docs/static/assets/sample.xlsx")
if err != nil {
    panic(err)
}

c, err := tbl.Col("Points")
if err != nil {
    panic(err)
}

fmt.Printf("name: %s\n", c.Name())
fmt.Printf("count: %v\n", c.Count())
```

Output:

```
name: Points
count: 15
```

---

## Behavior

- Searches the table's internal map for the specified column name.
- Duplicates the column's data slice to isolate the returned `Column` from side effects.
- Automatically infers whether the column is categorical based on its values (using a default threshold of 3 distinct unique values).
- Returns an error if:
  - The requested column name is not found in the table.

---

## See Also

### Related Types

- [`Column`](../../../column/) — represents a standalone column structure

### Related Methods

- [`Columns()`](../columns) — returns column names
- [`Len()`](../len) — returns the number of rows
- [`Display()`](../display) — prints the table
