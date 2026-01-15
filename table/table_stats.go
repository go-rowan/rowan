package table

import "fmt"

// Stats computes and displays descriptive statistics for all numeric columns in the table.
//
// The statistics include count, missing values, mean, standard deviation, minimum, quartiles (Q1, median, Q3), and maximum. Only columns containing numeric data are included in the output.
//
// The result is rendered as a table and printed directly to stdout.
// This method does not return a value and does not modify the original table.
// If the table is nil or empty, a message is printed instead.
func (t *Table) Stats() {
	if t == nil || t.Len() == 0 {
		fmt.Println("table is empty or nil")
	}

	statsColumns := []string{
		"Column", "Count", "Missing", "Mean", "Std", "Min", "Q1", "Median", "Q3", "Max",
	}
	statsData := make(map[string][]any, len(statsColumns))
	for _, c := range statsColumns {
		statsData[c] = []any{}
	}

	for _, c := range t.Columns() {
		col, err := t.Col(c)
		if err != nil {
			continue
		}

		if !isNumericColumn(col) {
			continue
		}

		mean, _ := col.Mean()
		std, _ := col.Std()
		min, _ := col.Min()
		q1, _ := col.Q1()
		median, _ := col.Median()
		q3, _ := col.Q3()
		max, _ := col.Max()

		statsData["Column"] = append(statsData["Column"], c)
		statsData["Count"] = append(statsData["Count"], col.Count())
		statsData["Missing"] = append(statsData["Missing"], col.Missing())
		statsData["Mean"] = append(statsData["Mean"], mean)
		statsData["Std"] = append(statsData["Std"], std)
		statsData["Min"] = append(statsData["Min"], min)
		statsData["Q1"] = append(statsData["Q1"], q1)
		statsData["Median"] = append(statsData["Median"], median)
		statsData["Q3"] = append(statsData["Q3"], q3)
		statsData["Max"] = append(statsData["Max"], max)
	}

	statsTable, err := New(statsData, statsColumns)
	if err != nil {
		fmt.Println("failed to create stats table:", err)
		return
	}

	statsTable.Display()
}
