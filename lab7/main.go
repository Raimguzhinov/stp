package main

import (
	"fmt"
	"stp/lab7/umemory"
)

func main() {
	// Пример использования с числами int.
	memory := umemory.NewTMemory(0)

	// Сохранение числа.
	memory.Store(10)
	fmt.Printf("State: %s, Number: %d\n", memory.GetState(), memory.GetNumber())

	// Добавление числа.
	if err := memory.Add(5, func(a, b int) (int, error) {
		return a + b, nil
	}); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("State: %s, Number: %d\n", memory.GetState(), memory.GetNumber())
	}

	// Очистка памяти.
	memory.Clear()
	fmt.Printf("State: %s, Number: %d\n", memory.GetState(), memory.GetNumber())
}
