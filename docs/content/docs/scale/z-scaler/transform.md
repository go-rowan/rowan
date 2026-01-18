---
title: "Transform()"
---

# Transform()

## Description
`Transform()` applies Z-score standardization to the specified columns of a table using statistics computed during Fit.

It returns a new table with standardized values while leaving the original table unchanged.  
Only numeric values are transformed; non-numeric values are preserved as-is.

Standardization is performed using the mean and standard deviation learned during the Fit phase.

## Signature

```go
func (s *ZScaler) Transform(t *table.Table, columns ...string) (*table.Table, error)
```

## Parameters

- ### t
  `*table.Table`  
  The input table to be transformed.

- ### columns
  `...string`  
  One or more column names to be standardized.

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
  - A column has zero standard deviation

If the transformation succeeds, the transformed table and `nil` error are returned.

## Behavior

- Creates a clone of the input table before applying transformations
- Standardizes numeric values using the formula:

```m
  (x - mean) / std
```

- Preserves non-numeric values without modification
- Applies standardization only to the specified columns
- Uses fitted columns by default when no columns are explicitly provided
- Returns an error if a column was not fitted
- Returns an error if a column has zero standard deviation
- Does not modify the original table

## Example Usage

After fitting a `ZScaler`, call `Transform()` to apply standardization to the desired columns.

`Transform()` can be called multiple times using the same fitted statistics.

## When to Use

Use `Transform()` when you need to:

- Standardize numeric features to zero mean and unit variance
- Apply consistent normalization to new or existing data
- Prepare data for statistical or machine learning algorithms
- Perform non-mutating transformations in data processing pipelines

## See Also
- [`(*ZScaler).Fit(t *table.Table, columns ...string)`](../fit)  
- `(*ZScaler).Mean()`  
- `(*ZScaler).Std()`  
- `(*ZScaler).IsFitted()`  
- `(*ZScaler).Reset()`
