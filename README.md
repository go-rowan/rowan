# Rowan

## Overview
Rowan is a Go library for working with tabular data. It provides a simple `Table` abstraction, along with common data transformation utilities such as column extracting, mapping, up to range and z-score scaling, to support data preprocessing workflows.

## Quick Start
The example below demonstrates creating a `Table`, applying a scaler, and exporting the result.

Download this excel file ([sample.xlsx](https://raw.githubusercontent.com/go-rowan/rowan/main/docs/assets/sample.xlsx)) to follow the example!

```go
tbl, err := rowan.FromExcel("docs/assets/sample.xlsx")
if err != nil {
    panic(err)
}
tbl.Display()

newTbl := tbl.Categorize()
newTbl.Display()

zScaler := scale.NewZScaler()

if err := zScaler.Fit(newTbl, "Score", "Points"); err != nil {
    panic(err)
}

finalTbl, err := zScaler.Transform(newTbl)
if err != nil {
    panic(err)
}
finalTbl.Display()

if err := finalTbl.WriteCSV("my_final_table.csv"); err != nil {
    panic(err)
}
```

## Features
- **Tabular data model**<br>
    A simple `Table` structure with named columns and consistent row length.
- **Data Ingestion**<br>
    Ingest tabular data from CSV, Excel, Google Sheets, and Go structs into a unified `Table`.
- **Table inspection and access**<br>
    Retrieve column names, table length, and column values.
- **Column operations**<br>
    Select, map, categorize, and transform column values for downstream processing.
- **Data transformation**<br>
    Built-in scalers such as `RangeScaler` and `ZScaler` with explicit Fit and Transform steps.
- **Data export**<br>
    Export `Table` data to common tabular formats such as CSV.

## Further Reading
For detailed documentation and additional examples, see the full documentation:
- [Rowan Documentation](https://go-rowan.github.io/rowan/)

## License
This project is licensed under the MIT License.  
See the [LICENSE](LICENSE) file for details.
