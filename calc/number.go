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
		return NewFractionNumber(num, den)
	}

	if strings.Contains(s, ".") {
		floatVal, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, errors.New("не удалось разобрать десятичное число")
		}
		str := strings.TrimPrefix(s, "-")
		parts := strings.Split(str, ".")
		decimalPlaces := len(parts[1])
		multiplier := int64(1)
		for i := 0; i < decimalPlaces; i++ {
			multiplier *= 10
		}
		intVal := int64(floatVal * float64(multiplier))
		if floatVal < 0 {
			intVal = -intVal
		}
		return NewFractionNumber(intVal, multiplier)
	}

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return NewFractionNumber(n, 1)
}
