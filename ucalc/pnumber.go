package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TPNumber struct {
	value     float64
	base      int
	precision int
}

func NewTPNumberFromString(val, baseStr, precStr string) (*TPNumber, error) {
	base, err1 := strconv.Atoi(baseStr)
	prec, err2 := strconv.Atoi(precStr)
	if err1 != nil || err2 != nil || base < 2 || base > 16 || prec < 0 {
		return nil, errors.New("некорректное основание или точность")
	}

	value, err := parseTPString(val, base)
	if err != nil {
		return nil, err
	}

	return &TPNumber{
		value:     value,
		base:      base,
		precision: prec,
	}, nil
}

func NewTPNumberFromNumber(val float64, base, precision int) *TPNumber {
	return &TPNumber{
		value:     val,
		base:      base,
		precision: precision,
	}
}

func (n *TPNumber) Add(other Number) Number {
	return &TPNumber{
		value:     n.value + other.(*TPNumber).value,
		base:      n.base,
		precision: n.precision,
	}
}

func (n *TPNumber) Sub(other Number) Number {
	return &TPNumber{
		value:     n.value - other.(*TPNumber).value,
		base:      n.base,
		precision: n.precision,
	}
}

func (n *TPNumber) Mul(other Number) Number {
	return &TPNumber{
		value:     n.value * other.(*TPNumber).value,
		base:      n.base,
		precision: n.precision,
	}
}

func (n *TPNumber) Div(other Number) (Number, error) {
	divisor := other.(*TPNumber).value
	if divisor == 0 {
		return nil, errors.New("деление на ноль")
	}
	return &TPNumber{
		value:     n.value / divisor,
		base:      n.base,
		precision: n.precision,
	}, nil
}

func (n *TPNumber) Square() Number {
	return &TPNumber{
		value:     n.value * n.value,
		base:      n.base,
		precision: n.precision,
	}
}

func (n *TPNumber) Inverse() (Number, error) {
	if n.value == 0 {
		return nil, errors.New("обращение нуля")
	}
	return &TPNumber{
		value:     1 / n.value,
		base:      n.base,
		precision: n.precision,
	}, nil
}

func (n *TPNumber) String() string {
	return formatTPString(n.value, n.base, n.precision)
}

func (n *TPNumber) Float64() float64 {
	return n.value
}

func (n *TPNumber) Copy() Number {
	return &TPNumber{
		value:     n.value,
		base:      n.base,
		precision: n.precision,
	}
}

// ========== Вспомогательные функции ==========

// parseTPString превращает строку вида "A.BC" в float64 в p-ичной системе
func parseTPString(s string, base int) (float64, error) {
	s = strings.ToUpper(s)
	parts := strings.SplitN(s, ".", 2)

	whole := 0
	for i := 0; i < len(parts[0]); i++ {
		d := digitToInt(rune(parts[0][i]))
		if d >= base || d == -1 {
			return 0, errors.New("цифра вне диапазона основания")
		}
		whole = whole*base + d
	}

	var frac float64 = 0
	if len(parts) == 2 {
		for i := len(parts[1]) - 1; i >= 0; i-- {
			d := digitToInt(rune(parts[1][i]))
			if d >= base || d == -1 {
				return 0, errors.New("цифра вне диапазона основания")
			}
			frac = (frac + float64(d)) / float64(base)
		}
	}

	return float64(whole) + frac, nil
}

func formatTPString(val float64, base int, prec int) string {
	if val < 0 {
		return "-" + formatTPString(-val, base, prec)
	}

	intPart := int64(val)
	frac := val - float64(intPart)

	var intStr string
	if intPart == 0 {
		intStr = "0"
	} else {
		for intPart > 0 {
			intStr = string(intToDigit(int(intPart%int64(base)))) + intStr
			intPart /= int64(base)
		}
	}

	if prec == 0 {
		return intStr
	}

	var fracStr string
	for i := 0; i < prec; i++ {
		frac *= float64(base)
		d := int(frac)
		frac -= float64(d)
		fracStr += string(intToDigit(d))
	}

	return fmt.Sprintf("%s.%s", intStr, fracStr)
}

func digitToInt(r rune) int {
	switch {
	case r >= '0' && r <= '9':
		return int(r - '0')
	case r >= 'A' && r <= 'F':
		return int(r-'A') + 10
	default:
		return -1
	}
}

func intToDigit(i int) rune {
	if i < 10 {
		return rune('0' + i)
	}
	return rune('A' + (i - 10))
}
