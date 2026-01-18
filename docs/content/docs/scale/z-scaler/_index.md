---
title: "ZScaler"
---

# ZScaler

`ZScaler` performs Z-score standardization, also known as standard scaling.

It transforms numeric values in selected columns so that each feature has zero mean and unit variance, using statistics learned during a fitting phase.

`ZScaler` is designed to work with Rowan tables and follows a stateful workflow: statistics are learned during Fit and later applied during Transform.

## Overview

`ZScaler` computes the mean and standard deviation for each selected feature and applies the following transformation formula:

```m
scaled = (x - mean) / std
```

The scaler stores per-column statistics internally and applies transformations in a non-mutating manner, meaning the original table is never modified.

Only numeric columns are considered during fitting and transformation.

## Lifecycle

ZScaler follows a simple lifecycle:

1. Create a new scaler instance using `NewZScaler()`
2. Call `Fit()` to learn per-column mean and standard deviation
3. Call `Transform()` to standardize data using the learned statistics
4. Optionally call `Reset()` to clear the internal state

## Internal State

`ZScaler` maintains the following internal state:

- `features`  
  A slice of column names that were fitted.

- `min`  
  A map storing the mean value for each fitted column.

- `max`  
  A map storing the standard deviation for each fitted column.

This state is updated only during Fit and cleared when `Reset()` is called.

## Constructor

### `NewZScaler()`

`NewZScaler()` creates and initializes a new `ZScaler` instance.

The returned scaler starts with an empty internal state and must be fitted before it can be used to transform data.

#### Example

```go
import (
	"github.com/go-rowan/rowan"
	"github.com/go-rowan/rowan/scale"
)

func main() {
	tbl, _ := rowan.FromExcel("data.xlsx")

	zScaler := scale.NewZScaler()
}
```

## Available Methods

### Common Methods

- `(*ZScaler).Fit(t *Table, columns ...string)`  
- `(*ZScaler).Transform(t *Table, columns ...string)`  
- `(*ZScaler).Features()`  
- `(*ZScaler).IsFitted()`  
- `(*ZScaler).Reset()`  

### ZScaler-Specific Methods

- `(*ZScaler).Mean(column string)`  
- `(*ZScaler).Std(column string)`  

## When to Use

Use `ZScaler` when you need to:

- Standardize numeric features to zero mean and unit variance
- Prepare data for algorithms that assume normally distributed inputs
- Perform consistent standardization learnt from a reference dataset

`ZScaler` is commonly used in statistical modeling and machine learning workflows.

## See Also

- [`RangeScaler`](../range-scaler)  
- [`(*ZScaler).Fit(t *Table, columns ...string)`](fit)  
- [`(*ZScaler).Transform(t *Table, columns ...string)`](transform)  
- `(*ZScaler).Mean(column string)`
