package main

import "fmt"

// BinarySearch выполняет бинарный поиск элемента в отсортированном массиве.
// Возвращает индекс найденного элемента или -1, если элемент не найден.
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// FindMinimumElement находит минимальный элемент в массиве.
func FindMinimumElement() {
	arr := []int{5, 3, 7, 2, 8}
	min, index := arr[0], 0
	for i, val := range arr {
		if val < min {
			min, index = val, i
		}
	}
	fmt.Printf("Минимум: %d на индексе %d\n", min, index)
}

// BubbleSort выполняет сортировку массива методом пузырька.
func BubbleSort() {
	arr := []int{5, 3, 7, 2, 8}
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Отсортированный массив:", arr)
}

// BinarySearch выполняет бинарный поиск в массиве.
func BinarySearchTask() {
	arr := []int{2, 3, 5, 7, 8}
	target := 5
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			fmt.Printf("Элемент %d найден на индексе %d\n", target, mid)
			return
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("Элемент %d не найден\n", target)
}

// ReverseArray выполняет реверс массива.
func ReverseArray() {
	arr := []int{5, 3, 7, 2, 8}
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	fmt.Println("Реверсированный массив:", arr)
}

// CyclicShift выполняет циклический сдвиг массива влево.
func CyclicShift() {
	arr := []int{5, 3, 7, 2, 8}
	k := 2
	result := append(arr[k:], arr[:k]...)
	fmt.Println("Массив после сдвига:", result)
}

// ReplaceElements заменяет элементы массива.
func ReplaceElements() {
	arr := []int{5, 3, 7, 2, 8}
	target, replacement := 7, 0
	for i := range arr {
		if arr[i] == target {
			arr[i] = replacement
		}
	}
	fmt.Println("Измененный массив:", arr)
}
