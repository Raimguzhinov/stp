package main

import (
	"errors"
	"stp/lab5/upnumber"
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

	// TPNumber: value|base|precision
	if strings.Contains(s, "|") {
		parts := strings.Split(s, "|")
		if len(parts) != 3 {
			return nil, errors.New("неверный формат p-ичного числа")
		}
		tpn, err := upnumber.NewTPNumberFromString(parts[0], parts[1], parts[2])
		if err != nil {
			return nil, err
		}
		if isNegative {
			tpn, _ = upnumber.NewTPNumberFromNumber(-tpn.GetValue(), tpn.GetBase(), tpn.GetPrecision())
		}
		return NewTPNumberAdapter(tpn), nil
	}

	// Complex: a+ib or a-ib
	if strings.Contains(s, "i") {
		s = strings.ReplaceAll(s, "−", "-") // handle minus
		s = strings.ReplaceAll(s, "−i", "-i")
		s = strings.ReplaceAll(s, "+i", "+i") // ensure delimiter
		realPart := "0"
		imagPart := "0"

		if strings.Contains(s, "+i") {
			parts := strings.Split(s, "+i")
			if len(parts) == 2 {
				realPart = parts[0]
				imagPart = parts[1]
			}
		} else if strings.Contains(s, "-i") {
			parts := strings.Split(s, "-i")
			if len(parts) == 2 {
				realPart = parts[0]
				imagPart = "-" + parts[1]
			}
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

	// Fraction
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

	// Decimal
	if strings.Contains(s, ".") {
		parts := strings.Split(s, ".")
		if len(parts) != 2 {
			return nil, errors.New("не удалось разобрать десятичное число")
		}
		numeratorStr := parts[0] + parts[1]
		numerator, err := strconv.ParseInt(numeratorStr, 10, 64)
		if err != nil {
			return nil, err
		}
		if isNegative {
			numerator = -numerator
		}
		denominator := int64(1)
		for range parts[1] {
			denominator *= 10
		}
		return NewFractionNumber(numerator, denominator)
	}

	// Integer
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	if isNegative {
		n = -n
	}
	return NewFractionNumber(n, 1)
}
