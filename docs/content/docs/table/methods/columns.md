---
title: "Columns()"
---

# Columns()

## Description

`Columns()` returns the list of column names contained in the table.

The order of the returned columns follows the table’s internal column order.

---

## Signature

```go
func (t *Table) Columns() []string
```

---

## Parameters

This method does not accept any parameters.

---

## Return Value

- `[]string`  
  A slice containing column names.

---

## Behavior

- Returns a copy of the column name slice
- Column order is deterministic
- Does not modify the original table
- If the table is empty, an empty slice is returned

---

## Example Usage

```go
table, _ := rowan.FromStructs([]struct {
    Name string
    Age  int
}{
    {"Alice", 30},
    {"Bob", 25},
})

cols := table.Columns()
fmt.Println(cols)
```

Output:

```
[Name Age]
```

---

## When to Use

Use `Columns()` when you need to:

- Inspect the structure of a table dynamically
- Validate column existence before calling Select, Drop, or MapCol
- Display table metadata to users

---

## Related Methods

- `Len()`
- `Col(name string)`
- `Select(cols ...string)`
- `Overview()`
