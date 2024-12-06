package main

import (
	"errors"
	"strconv"
)

type NumberOperations struct{}

func (n NumberOperations) OddDigitsReversed(a int) int {
	result := 0
	multiplier := 1

	str := strconv.Itoa(a)
	for i := 0; i < len(str); i += 2 {
		digit := int(str[i] - '0')
		result += digit * multiplier
		multiplier *= 10
	}

	return result
}

func (n NumberOperations) MaxEvenDigitPosition(a int) (int, error) {
	maxDigit := -1
	position := -1
	currentPos := 1

	str := strconv.Itoa(a)
	for i := 1; i < len(str); i += 2 {
		digit := int(str[i] - '0')
		if digit%2 == 0 && digit > maxDigit {
			maxDigit = digit
			position = currentPos
		}
		currentPos++
	}

	if maxDigit == -1 {
		return -1, errors.New("no even digits found")
	}

	return position, nil
}

func (n NumberOperations) RotateDigits(a, positions int) int {
	str := strconv.Itoa(a)
	length := len(str)

	shift := positions % length
	shiftedStr := str[length-shift:] + str[:length-shift]

	result, _ := strconv.Atoi(shiftedStr)
	return result
}

func (n NumberOperations) SumAboveAntiDiagonal(A [][]float64) (float64, error) {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0, errors.New("array is empty")
	}

	sum := 0.0
	length := len(A)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if int(A[i][j])%2 == 0 {
				sum += A[i][j]
			}
		}
	}

	return sum, nil
}
