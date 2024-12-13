package main

import "fmt"

func main() {
	tasks := []struct {
		Name string
		Func func()
	}{
		{"Find Minimum Element", FindMinimumElement},
		{"Bubble Sort", BubbleSort},
		{"Binary Search", BinarySearchTask},
		//{"Find Minimum in 2D Array", FindMinimumElement},
		{"Reverse Array", ReverseArray},
		{"Cyclic Shift", CyclicShift},
		{"Replace Elements", ReplaceElements},
	}

	var results []Metrics

	// Выполняем каждую задачу через универсальный обработчик
	for _, task := range tasks {
		result := processTask(task.Name, task.Func)
		results = append(results, result)
	}

	// Печать итоговой таблицы
	fmt.Printf("%-25s %-5s %-5s %-5s %-5s %-5s %-5s %-10s %-8s %-7s %-8s %-8s %-7s %-7s %-7s\n",
		"Task", "η₁", "η₂", "η", "N₁", "N₂", "N", "N^", "V*", "V", "L", "I", "T^1", "T^2", "T^3")
	for _, r := range results {
		fmt.Printf("%-25s %-5d %-5d %-5d %-5d %-5d %-5d %-10.2f %-8.2f %-7.2f %-8.2f %-8.2f %-7.2f %-7.2f %-7.2f\n",
			r.TaskName, r.Operators, r.Operands, r.DictionaryLength,
			r.OperatorCount, r.OperandCount, r.ImplementationLength,
			r.PredictedLength, r.PotentialVolume, r.RealVolume,
			r.ProgramLevel, r.IntellectContent, r.PredictedTime1,
			r.PredictedTime2, r.PredictedTime3)
	}
}
