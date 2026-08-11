---
title: "Display()"
---

# Display()

## Description

`Display()` prints the table to standard output (`stdout`) in a table layout.

It calculates the necessary column widths based on headers and row data, centers column headers, and formats values according to their types (e.g., center-aligning numeric and boolean values, and left-aligning text values).

---

## Signature

```go
func (t *Table) Display()
```

---

## Parameters

This method does not accept any parameters, at least for now.

---

## Return Values

None. Output is printed directly to standard output.

---

## Example

```
tbl, err := rowan.FromExcel("docs/static/assets/sample.xlsx")
if err != nil {
  panic(err)
}
tbl.Display()
```

Output:

```
--------------------------------------------
|      Name      | Gender | Score | Points |
--------------------------------------------
| Andi Pratama   | Male   | 82.50 |   12   |
| Siti Aisyah    | Female | 90.20 |   18   |
| Budi Santoso   | Male   | 74.80 |   8    |
| Rina Oktaviani | Female | 88.00 |   16   |
| Dimas Saputra  | Male   | 91.30 |   19   |
| Rizky Maulana  | Male   | 80.40 |   13   |
| Nabila Putri   | Female | 85.60 |   14   |
| Ahmad Fauzan   | Male   | 78.90 |   11   |
| Dewi Lestari   | Female | 92.10 |   19   |
| Ayu Wulandari  | Female | 89.50 |   17   |
| Putri Maharani | Female | 87.70 |   14   |
| Fajar Nugroho  | Male   | 76.20 |   9    |
| Ilham Ramadhan | Male   | 83.10 |   12   |
| Clara Novita   | Female | 86.80 |   14   |
| Yoga Prakoso   | Male   | 79.00 |   8    |
--------------------------------------------
```

---

## Behavior

- Renders all rows and columns using the table's current column order.
- Dynamically adjusts column widths to accommodate the longest string representation of values or header names.
- Alignments and formatting rules applied:
  - Header titles are centered.
  - Numeric values (integers, floats) and booleans are centered. Float values are formatted to two decimal places (`%.2f`).
  - Text and other non-numeric values are left-aligned.
- Prints `"nil"` if the table pointer is `nil`.
- Prints `"-- empty --"` if the table contains zero rows.

---

## Related Methods

- [`Overview()`](../overview) — prints summary of the table
- `Stats()`
- [`New()`](../../creation/new) — constructs a `Table` from a map
- [`FromExcel()`](../../creation/from-excel) — constructs a `Table` from an Excel file