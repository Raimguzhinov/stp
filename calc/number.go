package main

import (
	"errors"
	"strconv"
	"strings"
)

type Number interface {
	Add(Number) Number
	Sub(Number) Number
	Mul(Number) Number
	Div(Number) (Number, error)
	Square() Number
	Inverse() (Number, error)
	String() string
	Float64() float64
	Copy() Number
}

func ParseNumber(s string) (Number, error) {
	s = strings.ReplaceAll(s, " ", "")
	isNegative := strings.HasPrefix(s, "-")
	if isNegative {
		s = strings.TrimPrefix(s, "-")
	}

	if strings.Contains(s, "/") {
		parts := strings.Split(s, "/")
		if len(parts) != 2 {
			return nil, errors.New("неверный формат дроби")
		}
		num, err1 := strconv.ParseInt(parts[0], 10, 64)
		den, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
			return nil, errors.New("не удалось разобрать числитель или знаменатель")
		}
		if isNegative {
			num = -num
		}
		return NewFractionNumber(num, den)
	}

	if strings.Contains(s, ".") {
		parts := strings.Split(s, ".")
		if len(parts) != 2 {
			return nil, errors.New("не удалось разобрать десятичное число")
		}
		intPart := parts[0]
		fracPart := parts[1]
		if intPart == "" || fracPart == "" {
			return nil, errors.New("неверный формат дроби")
		}

		numeratorStr := intPart + fracPart
		numerator, err := strconv.ParseInt(numeratorStr, 10, 64)
		if err != nil {
			return nil, err
		}
		if isNegative {
			numerator = -numerator
		}
		denominator := int64(1)
		for i := 0; i < len(fracPart); i++ {
			denominator *= 10
		}
		return NewFractionNumber(numerator, denominator)
	}

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	if isNegative {
		n = -n
	}
	return NewFractionNumber(n, 1)
}
