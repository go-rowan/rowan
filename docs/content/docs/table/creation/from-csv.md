---
title: "FromCSV()"
---

# FromCSV()

## Description

`FromCSV()` creates a `Table` by reading data from a CSV file.

This function is the primary entry point when you want to load tabular data from a CSV source into Rowan. It reads the file, parses its rows and columns, and converts them into a `Table` structure.

---

## Signature

```go
FromCSV(path string, opts ...CSVOption) (*Table, error)
```

---

## Parameters

- `path`  
  The file path to the CSV file.

- `opts`  
  Optional CSV configuration options. These options allow you to customize how the CSV file is parsed, such as changing the delimiter.

---

## Return Values

- `*Table`  
  A pointer to the resulting `Table` containing the CSV data.

- `error`  
  An error is returned if the CSV file cannot be read or parsed.

---

## Behavior

- Reads the CSV file from the given path.
- Uses the provided CSV options (if any) to configure parsing behavior.
- Automatically infers column names from the CSV header.
- Returns an error if:
  - The file cannot be read.
  - The CSV format is invalid.
  - The data cannot be converted into a `Table`.

---

## Notes

- The returned `Table` is immutable; operations like `Select()`, `Drop()`, or `AddColumn()` return new tables.
- Column values are stored as []any and may require type conversion for numeric operations.
- CSV parsing behavior can be customized using helper options such as `WithDelimiter()`.

---

## Related Functions

- `WithDelimiter()`  
  Customize the CSV delimiter character.

- `FromStructs()`  
  Create a `Table` from a slice of structs instead of a CSV file.

---

## Example Usage

See this example for practical usage of `FromCSV()`.

Assume we have a CSV file named "data.csv" as below:

```
Name,Age,Score,Active
Alice,21,88.5,true
Bob,20,92,false
Charlie,18,79.25,true
```

The following code will read the file and construct a `Table`

```go
tbl, _ := rowan.FromCSV("data.csv")

tbl.First().Display()
tbl.Overview()
```

Output:

```
----------------------------------
|  Name   | Age | Score | Active |
----------------------------------
| Alice   | 21  | 88.50 |  true  |
| Bob     | 20  | 92.00 | false  |
| Charlie | 18  | 79.25 |  true  |
----------------------------------

Table Overview
Rows: 3
Columns:
-------------------
|  Name  |  Type  |
-------------------
| Name   | string |
| Age    | int    |
| Score  | float  |
| Active | bool   |
-------------------
```

## See Also

- [`FromExcel(path string, argOpts ...ExcelOption)`](../from-excel) — constructs a `Table` from an Excel file
- [`FromStructs(rows []T)`](../from-structs) — constructs a `Table` from a slice of structs
- `Overview()`
- `Display()`
