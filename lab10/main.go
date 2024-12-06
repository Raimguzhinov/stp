package main

import (
	"fmt"
	"stp/lab10/uset"
)

func main() {
	// Создание множеств
	set1 := uset.NewSet[int]()
	set2 := uset.NewSet[int]()

	set1.Add(1)
	set1.Add(2)
	set2.Add(2)
	set2.Add(3)

	// Объединение
	union := set1.Union(set2)
	fmt.Println("Union:", union.Elements()) // Output: [1 2 3]

	// Разность
	diff := set1.Difference(set2)
	fmt.Println("Difference:", diff.Elements()) // Output: [1]

	// Пересечение
	intersection := set1.Intersection(set2)
	fmt.Println("Intersection:", intersection.Elements()) // Output: [2]
}
