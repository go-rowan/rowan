---
title: "FromStructs()"
---

# FromStructs()

## Description

`FromStructs()` creates a `Table` from a slice of structs.

This constructor allows you to convert in-memory structured data into a `Table` by using reflection. Each exported field of the struct becomes a column, and each element in the slice becomes a row.

---

## Signature

```go
FromStructs[T any](rows []T) (*Table, error)
```

---

## Parameters

- `rows`  
  A slice of structs used as the data source for the `Table`.

---

## Struct Field Rules

- Only exported fields are included.
- Fields can be customized using struct tags.
- A custom column name can be provided via the &#96;rowan&#96; tag.
- Fields tagged with &#96;rowan:"-"&#96; are ignored.

---

## Return Values

- `*Table`  
  A pointer to the resulting `Table` constructed from the struct slice.

- `error`  
  An error is returned if the input slice is empty or does not contain structs.

---

## Behavior

- Inspects the struct type using reflection.
- Uses exported struct fields as columns.
- Iterates over each struct in the slice to populate table rows.
- Returns an error if:
  - The slice is empty.
  - The element type is not a struct.

---

## Notes

- Field order in the struct determines column order in the `Table`.
- All values are stored as `[]any` and may require type conversion for numeric operations.
- Struct tags provide full control over column naming and inclusion.

---

## Example Usage

See this example for practical usage of `FromStructs()` and seeing how struct tags affect the resulting columns.

Assume we have `Member` struct as below:

```go
type Member struct {
  Name   string
  Age    int     `rowan:"-"`
  Score  float32 `rowan:"Average Point"`
  Active bool    `rowan:"Status"`
}
```

The following code demonstrates creating a `Table` from a slice of `Member` instances. 

```go
memberSlice := []Member{
  {Name: "Alice", Age: 21, Score: 88.5, Active: true},
  {Name: "Bob", Age: 20, Score: 92, Active: false},
  {Name: "Charlie", Age: 18, Score: 79.25, Active: true},
}

tbl, _ := rowan.FromStructs(memberSlice)

tbl.First().Display()
tbl.Overview()
```

Output:

```
------------------------------------
|  Name   | Average Point | Status |
------------------------------------
| Alice   |     88.50     |  true  |
| Bob     |     92.00     | false  |
| Charlie |     79.25     |  true  |
------------------------------------

Table Overview
Rows: 3
Columns:
--------------------------
|     Name      |  Type  |
--------------------------
| Name          | string |
| Average Point | float  |
| Status        | bool   |
--------------------------
```

## See Also

- `New(data map[string][]any, columnsOrder ...[]string)`
- [`FromCSV(path string, opts ...CSVOption)`](../from-csv) — constructs a `Table` from a CSV file
- `Overview()`
- `Display()`
