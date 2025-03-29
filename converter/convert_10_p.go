package main

import (
	"strings"
)

func Convert10ToP(value float64, base int, precision int) string {
	neg := value < 0
	if neg {
		value = -value
	}

	intPart := int(value)
	fracPart := value - float64(intPart)

	intStr := intToP(intPart, base)
	fracStr := fracToP(fracPart, base, precision)

	result := intStr
	if fracStr != "" {
		result += "." + fracStr
	}
	if neg {
		result = "-" + result
	}
	return strings.ToUpper(result)
}

func intToP(n int, base int) string {
	if n == 0 {
		return "0"
	}
	digits := ""
	for n > 0 {
		digits = string(intToChar(n%base)) + digits
		n /= base
	}
	return digits
}

func fracToP(f float64, base int, precision int) string {
	result := ""
	for i := 0; i < precision; i++ {
		f *= float64(base)
		digit := int(f)
		result += string(intToChar(digit))
		f -= float64(digit)
	}
	return result
}

func intToChar(n int) rune {
	if n >= 0 && n <= 9 {
		return rune('0' + n)
	}
	if n >= 10 && n <= 15 {
		return rune('A' + (n - 10))
	}
	return '?'
}
