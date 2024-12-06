package main

import (
	"errors"
	"fmt"
	"strings"
)

// Matrix - структура для работы с матрицами
type Matrix struct {
	data [][]int
	rows int
	cols int
}

// 1. Конструктор
func NewMatrix(data [][]int) (*Matrix, error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, errors.New("matrix dimensions must be greater than 0")
	}
	rows := len(data)
	cols := len(data[0])
	for _, row := range data {
		if len(row) != cols {
			return nil, errors.New("all rows must have the same number of columns")
		}
	}
	return &Matrix{data: data, rows: rows, cols: cols}, nil
}

// 2. Сложение матриц
func (m *Matrix) Add(b *Matrix) (*Matrix, error) {
	if m.rows != b.rows || m.cols != b.cols {
		return nil, errors.New("matrices dimensions must match")
	}

	result := make([][]int, m.rows)
	for i := range result {
		result[i] = make([]int, m.cols)
		for j := range result[i] {
			result[i][j] = m.data[i][j] + b.data[i][j]
		}
	}
	return NewMatrix(result)
}

// 3. Вычитание матриц
func (m *Matrix) Sub(b *Matrix) (*Matrix, error) {
	if m.rows != b.rows || m.cols != b.cols {
		return nil, errors.New("matrices dimensions must match")
	}

	result := make([][]int, m.rows)
	for i := range result {
		result[i] = make([]int, m.cols)
		for j := range result[i] {
			result[i][j] = m.data[i][j] - b.data[i][j]
		}
	}
	return NewMatrix(result)
}

// 4. Умножение матриц
func (m *Matrix) Mul(b *Matrix) (*Matrix, error) {
	if m.cols != b.rows {
		return nil, errors.New("matrices cannot be multiplied: incompatible dimensions")
	}

	result := make([][]int, m.rows)
	for i := range result {
		result[i] = make([]int, b.cols)
		for j := 0; j < b.cols; j++ {
			for k := 0; k < m.cols; k++ {
				result[i][j] += m.data[i][k] * b.data[k][j]
			}
		}
	}
	return NewMatrix(result)
}

// 5. Сравнение матриц
func (m *Matrix) Equals(b *Matrix) bool {
	if m.rows != b.rows || m.cols != b.cols {
		return false
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if m.data[i][j] != b.data[i][j] {
				return false
			}
		}
	}
	return true
}

// 6. Транспонирование
func (m *Matrix) Transpose() (*Matrix, error) {
	if m.rows != m.cols {
		return nil, errors.New("matrix must be square to transpose")
	}

	result := make([][]int, m.cols)
	for i := range result {
		result[i] = make([]int, m.rows)
		for j := range result[i] {
			result[i][j] = m.data[j][i]
		}
	}
	return NewMatrix(result)
}

// 7. Минимальный элемент
func (m *Matrix) Min() int {
	min := m.data[0][0]
	for i := range m.data {
		for j := range m.data[i] {
			if m.data[i][j] < min {
				min = m.data[i][j]
			}
		}
	}
	return min
}

// 8. Преобразование в строку
func (m *Matrix) ToString() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, row := range m.data {
		sb.WriteString(fmt.Sprintf("{%v}", row))
		if i < m.rows-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")
	return sb.String()
}

// 9. Получить элемент с индексами i,j
func (m *Matrix) Get(i, j int) (int, error) {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return 0, errors.New("index out of bounds")
	}
	return m.data[i][j], nil
}

// 10. Взять число строк
func (m *Matrix) Rows() int {
	return m.rows
}

// 11. Взять число столбцов
func (m *Matrix) Cols() int {
	return m.cols
}

func main() {
	// Пример использования
	data1 := [][]int{{1, 2}, {3, 4}}
	data2 := [][]int{{5, 6}, {7, 8}}

	matrix1, _ := NewMatrix(data1)
	matrix2, _ := NewMatrix(data2)

	// Сложение
	sumMatrix, _ := matrix1.Add(matrix2)
	fmt.Println("Сумма матриц:", sumMatrix.ToString())

	// Вычитание
	subMatrix, _ := matrix1.Sub(matrix2)
	fmt.Println("Разность матриц:", subMatrix.ToString())

	// Умножение
	mulMatrix, _ := matrix1.Mul(matrix2)
	fmt.Println("Произведение матриц:", mulMatrix.ToString())

	// Транспонирование
	transposedMatrix, _ := matrix1.Transpose()
	fmt.Println("Транспонированная матрица:", transposedMatrix.ToString())

	// Минимальный элемент
	fmt.Println("Минимальный элемент:", matrix1.Min())

	// Получение элемента
	element, _ := matrix1.Get(0, 1)
	fmt.Println("Элемент [0,1]:", element)

	// Число строк и столбцов
	fmt.Println("Число строк:", matrix1.Rows())
	fmt.Println("Число столбцов:", matrix1.Cols())
}
