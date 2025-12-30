package table

import "fmt"

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
