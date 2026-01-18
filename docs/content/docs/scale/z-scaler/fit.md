---
title: "Fit()"
---

# Fit()

## Description
`Fit()` computes the mean and standard deviation for the specified columns from the provided table.

The computed statistics are stored internally and later used during Transform to perform Z-score standardization.

`Fit()` overwrites any previously stored statistics for the same columns.

## Signature

```go
func (s *ZScaler) Fit(t *table.Table, columns ...string) error
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
  - A column has zero standard deviation

  If all columns are fitted successfully, `nil` is returned.

## Behavior
- Computes mean and standard deviation for each specified column
- Considers only numeric values when computing statistics
- Ignores non-numeric values during computation
- Stores computed statistics internally
- Appends fitted column names to the feature list
- Overwrites previously fitted statistics for the same columns
- Does not modify the input table

## Example Usage
After creating a `ZScaler` instance, call `Fit()` with a table and the columns to standardize.

Once `Fit()` succeeds, the scaler is ready to be used with `Transform()`.

## When to Use
Use `Fit()` when you need to:

- Learn per-column mean and standard deviation from a reference dataset
- Prepare a `ZScaler` for consistent standardization
- Refit the scaler with updated or different training data

## See Also
- [`(*ZScaler).Transform(t *table.Table, columns ...string)`](../transform)  
- `(*ZScaler).Mean(column string)`  
- `(*ZScaler).Std(column string)`  
- `(*ZScaler).IsFitted()`  
- `(*ZScaler).Reset()`
