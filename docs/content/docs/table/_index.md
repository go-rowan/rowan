---
title: "Table"
---

# Table

## Description

`Table` is the primary in-memory tabular data structure in Rowan, responsible for managing two-dimensional data arranged by named columns and ordered rows.

It maintains explicit column ordering, a map of underlying column data slices, and the total row length of the dataset. `Table` serves as the central hub for data manipulation, IO operations, row/column extractions, and summary statistics.

---

## Type Definition

```go
type Table struct {
	columns []string
	data    map[string][]any
	length  int
}
```

---

## Struct Fields

- `columns`  
  A slice of strings (`[]string`) preserving the explicit sequence and order of columns in the table.

- `data`  
  A map (`map[string][]any`) where each key corresponds to a column name and its value holds the slice of row data for that column.

- `length`  
  An integer representing the total number of rows across all columns in the table.

---

## Behavior & Properties

- Enforces strict row length consistency across all contained columns.
- Preserves predictable column sequences independent of Go's non-deterministic map iteration via its internal `columns` order slice.
- Provides thread-safe isolated copies when extracting individual series via methods like `Col()`.

---

## Available Methods

### Constructors & Factories

- [`New()`](creation/new) — constructs a `Table` from a map
- [`FromStructs()`](creation/from-structs) — constructs a `Table` from a slice of structs
- [`FromCSV()`](creation/from-csv) — constructs a `Table` from a CSV file

### Instance Methods

- [`Len()`](methods/len) — returns the number of rows
- [`Columns()`](methods/columns) — returns column names
- [`Display()`](methods/display) — prints the table


---

## See Also

### Related Types

- [`Column`](../column/) — represents a standalone column structure

### Related Methods

- [`Col()`](methods/col) — extracts a `Column` instance from a `Table`
