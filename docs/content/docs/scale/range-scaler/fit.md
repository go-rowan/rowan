---
title: "Fit()"
---

# Fit()

## Description
`Fit()` computes and stores the minimum and maximum values for the specified columns from the provided table.

The computed statistics are stored internally and later used during Transform to perform range-based scaling.

`Fit()` overwrites any previously stored statistics for the same columns.

## Signature

```go
func (s *RangeScaler) Fit(t *table.Table, columns ...string) error
```

## Parameters

- ### t
  `*table.Table`  
  The input table used to compute column statistics.

- ### columns
  `...string`  
  Zero or more column names to be fitted.

  Only the specified columns are processed.

  Calling `Fit()` with zero column names results in a no-op. In this case, `IsFitted()` remains returning `false`.

## Return Value
- ### `error`  

  `Fit()` returns an error if:

  - The provided table is `nil`
  - A specified column does not exist
  - A column contains no numeric values

  If all columns are fitted successfully, `nil` is returned.

## Behavior
- Computes minimum and maximum values for each specified column
- Considers only numeric values when computing statistics
- Ignores non-numeric values during computation
- Stores computed statistics internally
- Appends fitted column names to the feature list
- Overwrites previously fitted statistics for the same columns
- Does not modify the input table
- If no column names are provided, `Fit()` performs no operation and the scaler remains unfitted

## Example Usage
After creating a `RangeScaler` instance, call `Fit()` with a table and the columns to be scaled.

Once `Fit()` succeeds, the scaler is ready to be used with `Transform()`.

## When to Use
Use `Fit()` when you need to:

- Learn per-column minimum and maximum values from a reference dataset
- Prepare a `RangeScaler` for consistent range-based normalization
- Refit the scaler with updated or different training data

## See Also
- [`(*RangeScaler).Transform(t *table.Table, columns ...string)`](../transform)  
- `(*RangeScaler).Min(column string)`  
- `(*RangeScaler).Max(column string)`  
- `(*RangeScaler).IsFitted()`  
- `(*RangeScaler).Reset()`
