---
title: "RangeScaler"
---

# RangeScaler

`RangeScaler` implements range-based feature scaling, commonly known as Min-Max scaling.

It scales numeric values in specified columns into a normalized range, typically between 0 and 1, using statistics learned during a fitting phase.

`RangeScaler` is designed to work with Rowan tables and follows a stateful workflow: statistics are learned during Fit and later applied during Transform.

## Overview

`RangeScaler` computes the minimum and maximum value for each selected feature and applies the following transformation formula:

```m
scaled = (x - min) / (max - min)
```

The scaler stores per-column statistics internally and applies transformations in a non-mutating manner, meaning the original table is never modified.

Only numeric columns are considered during fitting and transformation.

## Lifecycle

RangeScaler follows a simple lifecycle:

1. Create a new scaler instance using `NewRangeScaler()`
2. Call `Fit()` to learn per-column minimum and maximum values
3. Call `Transform()` to scale data using the learned statistics
4. Optionally call `Reset()` to clear the internal state

## Internal State

`RangeScaler` maintains the following internal state:

- `features`  
  A slice of column names that were fitted.

- `min`  
  A map storing the minimum value for each fitted column.

- `max`  
  A map storing the maximum value for each fitted column.

This state is updated only during Fit and cleared when `Reset()` is called.

## Constructor

### `NewRangeScaler()`

`NewRangeScaler()` creates and initializes a new `RangeScaler` instance.

The returned scaler starts with an empty internal state and must be fitted before it can be used to transform data.

#### Example

```go
import (
	"github.com/go-rowan/rowan"
	"github.com/go-rowan/rowan/scale"
)

func main() {
	tbl, _ := rowan.FromExcel("data.xlsx")

	rangeScaler := scale.NewRangeScaler()
}
```

## Available Methods

### Common Methods

- `(*RangeScaler).Fit(t *Table, columns ...string)`  
- `(*RangeScaler).Transform(t *Table, columns ...string)`  
- `(*RangeScaler).Features()`  
- `(*RangeScaler).IsFitted()`  
- `(*RangeScaler).Reset()`  

### RangeScaler-Specific Methods

- `(*RangeScaler).Min(column string)`  
- `(*RangeScaler).Max(column string)`  

## When to Use

Use `RangeScaler` when you need to:

- Normalize numeric features into a fixed range
- Prepare data for algorithms sensitive to feature scale
- Apply consistent scaling learned from a reference dataset
- Perform min-max normalization in a reproducible way

`RangeScaler` is especially useful when feature bounds are known or when preserving relative distances within a fixed range is important.

## See Also

- `ZScaler`  
- `(*RangeScaler).Fit(t *Table, columns ...string)`  
- `(*RangeScaler).Transform(t *Table, columns ...string)`  
- `(*RangeScaler).Min(column string)`
