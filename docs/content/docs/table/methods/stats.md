---
title: "Stats()"
---

# Stats()

## Description

`Stats()` computes summary descriptive statistics for every numeric column in the table and outputs the results as a formatted table to standard output.

It automatically filters out non-numeric columns, calculates key statistical metrics for the remaining columns, constructs an internal summary table, and renders it using `Display()`.

---

## Signature

```go
func (t *Table) Stats()
```

---

## Parameters

None.

---

## Return Values

None. The calculated summary table is printed directly to standard output.

---

## Example

```go
tbl, err := rowan.FromExcel("docs/static/assets/sample.xlsx")
if err != nil {
  panic(err)
}
tbl.Stats()
```

Output:

```
------------------------------------------------------------------------------------
| Column | Count | Missing | Mean  | Std  |  Min  |  Q1   | Median |  Q3   |  Max  |
------------------------------------------------------------------------------------
| Score  |  15   |    0    | 84.41 | 5.59 | 74.80 | 79.70 | 85.60  | 88.75 | 92.10 |
| Points |  15   |    0    | 13.60 | 3.70 | 8.00  | 11.50 | 14.00  | 16.50 | 19.00 |
------------------------------------------------------------------------------------
```

---

## Behavior

- Iterates through all columns in the table and filters for numeric types (`isNumericColumn`). Non-numeric columns are skipped.
- Computes the following metrics for each numeric column:
  - `Count`: Number of present values.
  - `Missing`: Number of missing / `nil` values.
  - `Mean`: Average value.
  - `Std`: Standard deviation.
  - `Min`: Minimum value.
  - `Q1`: First quartile (25th percentile).
  - `Median`: Median value (50th percentile).
  - `Q3`: Third quartile (75th percentile).
  - `Max`: Maximum value.
- Renders the resulting statistics using `Display()`, maintaining a fixed column layout: `Column`, `Count`, `Missing`, `Mean`, `Std`, `Min`, `Q1`, `Median`, `Q3`, `Max`.
- Does not modify the underlying data of the original `Table`.
- Prints `"table is empty or nil"` if the table pointer is `nil` or has no rows (`t.Len() == 0`).

---

## Related Methods

- [`Display()`](../display) — prints the table
- [`Overview()`](../overview) — prints summary of the table
- [`New()`](../../creation/new) — constructs a `Table` from a map