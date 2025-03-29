package main

import (
	"math"
	"strings"
)

func ConvertPTo10(input string, base int) float64 {
	neg := false
	if strings.HasPrefix(input, "-") {
		neg = true
		input = input[1:]
	}

	parts := strings.Split(input, ".")
	intPart := parts[0]
	fracPart := ""
	if len(parts) == 2 {
		fracPart = parts[1]
	}

	value := 0.0
	// целая часть
	for i, c := range intPart {
		d := charToInt(c)
		pow := float64(len(intPart) - i - 1)
		value += float64(d) * math.Pow(float64(base), pow)
	}

	// дробная часть
	for i, c := range fracPart {
		d := charToInt(c)
		pow := float64(-(i + 1))
		value += float64(d) * math.Pow(float64(base), pow)
	}

	if neg {
		value = -value
	}
	return value
}

func charToInt(c rune) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')
	case c >= 'A' && c <= 'F':
		return int(c - 'A' + 10)
	case c >= 'a' && c <= 'f':
		return int(c - 'a' + 10)
	default:
		return -1
	}
}
