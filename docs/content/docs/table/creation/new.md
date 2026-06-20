---
title: "New()"
---

# New()

## Description

`New()` constructs a new `Table` from a map of column names to slices of values.

This function serves as a primary constructor to initialize a `Table` in Rowan using existing in-memory data structures, ensuring that all structural rules—such as consistent column lengths—are strictly validated.

---

## Signature

```go
func New(data map[string][]any, columnsOrder ...[]string) (*Table, error)
```

---

## Parameters

- ### data
  `map[string][]any`    
  A map where each key represents a column name and its associated value is a slice containing the rows for that column.

- ### `columnsOrder`
  `...[]string`    
  An optional, variadic slice specifying the desired sequence of columns in the Table. Only the first slice provided will be used. If omitted or nil, the final column order will follow the map's default iteration sequence, which is non-deterministic in Go.

---

## Return Values

- `*Table`  
  A pointer to the constructed Table instance populated with the supplied map data and determined column order.

- `error`  
  An error is returned if the configuration or data validation fails.

  It returns an error if:
  - The input data map is empty.
  - Column lengths are inconsistent across the dataset.
  - A column specified in columnsOrder does not exist in the map data.

---

## Behavior

- Validates that the input map is not empty.
- Checks and enforces that every column slice contains the exact same number of elements (consistent length).
- Applies explicit ordering to the table columns if a valid columnsOrder slice is supplied, validating that all requested columns exist within the data map.

## Example Usage

```go
myData := map[string][]any{
  "Name": {"Alice", "Bob", "Charlie"},
  "Age":  {28, 31, 30},
}

tbl, err := table.New(myData)
if err != nil {
  panic(err)
}

tbl.Display()
```

Output:

```
-----------------
|  Name   | Age |
-----------------
| Alice   | 28  |
| Bob     | 31  |
| Charlie | 30  |
-----------------
```

## Related Methods

- `(*Table).Display()`
- [`FromCSV(path string, opts ...CSVOption)`](../from-csv) — constructs a `Table` from a CSV file
- [`FromExcel(path string, argOpts ...ExcelOption)`](../from-excel) — constructs a `Table` from an Excel file
- [`FromStructs(rows []T)`](../from-structs) — constructs a `Table` from a slice of structs
