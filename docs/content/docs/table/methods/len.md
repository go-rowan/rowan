---
title: "Len()"
---

# Len()

## Description

`Len()` returns the number of rows in the table.

This method reflects how many records are stored in the table, not how many columns it has.

## Signature

```go
Len() int
```

## Parameters

This method does not accept any parameters.

---

## Return Value

- `int`  
  An integer number of rows value.

---

## Behavior

- Returns the total number of rows in the table.
- Returns 0 if the table contains no rows.
- The value is consistent across all columns.

## Notes

- A table can have zero rows but still have one or more columns.
- Len does not perform any computation; it simply returns the stored row count.
- This method does not modify the table.

## Example Usage

See this example for practical usage of `Len()`.

Assume we have a `Table` stored in the variable `tbl`.

```go
tbl.Display()
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
```

The following code demonstrates getting the information of table length (number of rows) and directly printing it.

```go
fmt.Println(tbl.Len())
```

Output:

```
3
```

## See Also

- [`Columns()`](../columns) — returns column names
- `First()` — returns the first N rows
- `Last()` — returns the last N rows
- `Sample()` — returns N randomly sampled rows
