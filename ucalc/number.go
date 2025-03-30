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

	switch currentMode {
	case ModeComplex:
		// Нормализуем строку
		s = strings.ReplaceAll(s, "−", "-")

		// Явно комплексный ввод
		if strings.Contains(s, "i") {
			s = strings.ReplaceAll(s, "i", "") // убираем i
			var realPart, imagPart string

			if strings.Contains(s, "+") {
				parts := strings.Split(s, "+")
				if len(parts) != 2 {
					return nil, errors.New("неверный формат комплексного числа (a+ib)")
				}
				realPart, imagPart = parts[0], parts[1]
			} else if len(s) >= 2 && strings.Contains(s[1:], "-") { // пропускаем возможный минус в начале
				idx := strings.LastIndex(s, "-")
				realPart, imagPart = s[:idx], s[idx:]
			} else {
				// ввод вида "5i"
				realPart = "0"
				imagPart = s
			}

			if realPart == "" {
				realPart = "0"
			}
			if imagPart == "" {
				imagPart = "1"
			} else if imagPart == "-" {
				imagPart = "-1"
			}

			r, err1 := strconv.ParseFloat(realPart, 64)
			i, err2 := strconv.ParseFloat(imagPart, 64)
			if err1 != nil || err2 != nil {
				return nil, errors.New("не удалось разобрать комплексное число")
			}
			if isNegative {
				r = -r
			}
			return NewComplexNumber(r, i), nil
		}

		// Автоконвертация вещественного числа в комплексное
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, errors.New("не удалось разобрать число")
		}
		if isNegative {
			val = -val
		}
		return NewComplexNumber(val, 0), nil

	case ModeTPNumber:
		tpn, err := NewTPNumberFromString(s, strconv.Itoa(currentBase), strconv.Itoa(currentPrecision))
		if err != nil {
			return nil, err
		}
		if isNegative {
			tpn = NewTPNumberFromNumber(-tpn.value, tpn.base, tpn.precision)
		}
		return tpn, nil

	case ModeFraction:
		// дробь
		if strings.Contains(s, "/") {
			parts := strings.Split(s, "/")
			if len(parts) != 2 {
				return nil, errors.New("неверный формат дроби")
			}
			num, err1 := strconv.ParseInt(parts[0], 10, 64)
			den, err2 := strconv.ParseInt(parts[1], 10, 64)
			if err1 != nil || err2 != nil {
				return nil, errors.New("не удалось разобрать дробь")
			}
			if isNegative {
				num = -num
			}
			return NewFractionNumber(num, den)
		}

		// десятичное
		if strings.Contains(s, ".") {
			parts := strings.Split(s, ".")
			if len(parts) != 2 {
				return nil, errors.New("неверный формат десятичного числа")
			}
			numStr := parts[0] + parts[1]
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				return nil, err
			}
			if isNegative {
				num = -num
			}
			den := int64(1)
			for range parts[1] {
				den *= 10
			}
			return NewFractionNumber(num, den)
		}

		// целое
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		if isNegative {
			n = -n
		}
		return NewFractionNumber(n, 1)
	}

	return nil, errors.New("не удалось определить тип числа")
}
