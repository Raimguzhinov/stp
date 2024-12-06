package main

import (
	"errors"
	"fmt"
	"math"
)

type MatrixOperations struct{}

func (m MatrixOperations) MinOfThree(a, b, c float64) float64 {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func (m MatrixOperations) SumEvenIndexElements(A [][]float64) (float64, error) {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0, errors.New("empty array")
	}

	sum := 0.0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if (i+j)%2 == 0 {
				sum += A[i][j]
			}
		}
	}
	return sum, nil
}

func (m MatrixOperations) MaxOnAndBelowDiagonal(A [][]float64) (float64, error) {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0, errors.New("empty array")
	}
	maxVal := math.Inf(-1)
	for i := 0; i < len(A); i++ {
		for j := 0; j <= i && j < len(A[i]); j++ {
			if A[i][j] > maxVal {
				maxVal = A[i][j]
			}
		}
	}
	return maxVal, nil
}

func main() {
	m := MatrixOperations{}
	fmt.Println(m.MinOfThree(1.5, 2.7, -0.2)) // Пример вызова первой функции
}
