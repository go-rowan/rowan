---
title: "Overview()"
---

# Overview()

## Description
`Overview()` prints a human-readable summary of the table to the standard output.

It displays the total number of rows and a column metadata table describing each column’s name and inferred data type.

This method is intended for quick inspection and debugging, not for programmatic access to metadata.

## Signature

```go
func (t *Table) Overview()
```

## Parameters
This method does not accept any parameters.

## Return Value
None.

`Overview()` does not return any value.  
All information is printed directly to standard output.

## Behavior
- Prints a table summary header followed by the total row count
- Displays column metadata including:
  - Column name
  - Inferred data type based on non-nil values
- Data type inference rules:
  - `int`, `int64` → "int"
  - `float32`, `float64` → "float"
  - `bool` → "bool"
  - `string` → "string"
- If all values in a column are `nil`, the type is reported as "unknown"
- Does not modify the original table
- If the receiver is `nil`, the string "nil" is printed and the method returns immediately

## Example Usage

```go
tbl, _ := rowan.FromExcel("data.xlsx")
tbl.Overview()
```

Output:

```
Table Overview
Rows: 15
Columns:
-------------------
|  Name  |  Type  |
-------------------
| Name   | string |
| Gender | string |
| Score  | float  |
| Points | int    |
-------------------
```

## When to Use
Use `Overview()` when you need to:

- Quickly inspect the structure and contents of a table during development
- Debug data ingestion results from CSV, Excel, Google Sheets, or structs
- Validate inferred column types before applying transformations

## Related Methods
- `Display()`
- `Col()`
- `Stats()`
- `Select(cols ...string)`
