---
title: "Column"
---

# Column

## Description

`Column` represents a single series or sequence of data within Rowan, containing its header name, raw element values, and inferred type metadata.

It is typically constructed and returned when extracting data from a `Table` (for instance, via `Col()`). A `Column` holds its own isolated data slice and supports various statistical, mathematical, and analytical operations.

---

## Type Definition

```go
type Column struct {
	name        string
	data        []any
	categorical bool
}
```

---

## Struct Fields

- `name`  
  The string identifier/header of the column.

- `data`  
  A slice (`[]any`) holding the raw row elements contained within the column.

- `categorical`  
  A boolean flag indicating whether the column values are treated as discrete categorical data (inferred based on unique value density).

---

## Behavior & Properties

- Operates as a standalone data structure isolated from its source `Table`.
- Stores elements dynamically using `[]any`, allowing flexible type handling across numeric, string, and missing values.
- Holds metadata used internally by analytical methods (such as categorical status detection).

---

## Available Methods

- `Count()` — Returns the number of non-nil elements in the column
- `Missing()` — Counts the number of missing/nil values
- `Mean()` — Computes the arithmetic mean for numeric columns
- `Std()` — Computes the standard deviation
- `Min()` / `Max()` — Finds the minimum and maximum values
- `Q1()` / `Median()` / `Q3()` — Calculates quartile statistical distributions
- many others...

---

## See Also

### Related Types
- [`Table`](../table/) — the primary multi-column tabular structure in Rowan

### Related Methods
- [`Col()`](../table/methods/col) — extracts a `Column` instance from a `Table`