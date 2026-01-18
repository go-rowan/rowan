---
title: "Transform()"
---

# Transform()

## Description
`Transform()` applies range-based scaling to the specified columns of a table.

It returns a new table with scaled values while leaving the original table unchanged.  
Only numeric values are transformed; non-numeric values are preserved as-is.

Scaling is performed using statistics learned during the Fit phase.

## Signature

```go
func (s *RangeScaler) Transform(t *table.Table, columns ...string) (*table.Table, error)
```

## Parameters

- ### t
  `*table.Table`  
  The input table to be transformed.

- ### columns
  `...string`  
  Zero or more column names to be scaled.

  If no column names are provided, `Transform()` uses the columns previously fitted by `Fit()`.

## Return Value

- ### `*table.Table`  
  A new table containing the transformed data.

- ### `error`  
  `Transform()` returns an error if:

  - The provided table is `nil`
  - A specified column does not exist
  - No columns are specified and the scaler has no fitted features
  - A column was not fitted
  - A column has zero range (min equals max)

If the transformation succeeds, the transformed table and `nil` error are returned.

## Behavior

- Creates a clone of the input table before applying transformations
- Scales numeric values using the formula:

```m
  (x - min) / (max - min)
```

- Preserves non-numeric values without modification
- Applies scaling only to the specified columns
- Uses fitted columns by default when no columns are explicitly provided
- Returns an error if a column was not fitted
- Returns an error if a column has zero range to avoid division by zero
- Does not modify the original table

## Example Usage
After fitting a `RangeScaler`, call `Transform()` to apply scaling to the desired columns.

`Transform()` can be called multiple times using the same fitted statistics.

## When to Use
Use `Transform()` when you need to:

- Apply min-max scaling to new or existing data
- Normalize numeric features into a fixed range
- Reuse scaling parameters learned from a reference dataset
- Perform non-mutating transformations in data processing pipelines

## See Also
- [`(*RangeScaler).Fit(t *table.Table, columns ...string)`](../fit)  
- `(*RangeScaler).Min()`  
- `(*RangeScaler).Max()`  
- `(*RangeScaler).IsFitted()`  
- `(*RangeScaler).Reset()`
