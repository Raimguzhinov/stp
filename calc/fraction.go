package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type FractionNumber struct {
	numerator   int64
	denominator int64
}

type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

var (
	ErrDivideByZero    = errors.New("ошибка деления на ноль")
	ErrZeroDenominator = errors.New("знаменатель не может равняться нулю")

	ZeroValue = &FractionNumber{
		numerator: 0, denominator: 1,
	}
	RevOneValue = &FractionNumber{
		numerator: -1, denominator: 1,
	}
	OneValue = &FractionNumber{
		numerator: 1, denominator: 1,
	}
)

func NewFractionNumber[T, K integer](numerator T, denominator K) (*FractionNumber, error) {
	if denominator == 0 {
		return ZeroValue, ErrZeroDenominator
	}
	if numerator == 0 {
		return ZeroValue, nil
	}

	n := int64(numerator)
	d := int64(denominator)
	if d < 0 {
		d *= -1
		n *= -1
	}
	gcf := gcd(abs(n), d)

	return &FractionNumber{
		numerator:   n / gcf,
		denominator: d / gcf,
	}, nil
}

func NewFractionFromString(s string) (*FractionNumber, error) {
	s = strings.ReplaceAll(s, " ", "")

	// Если строка содержит "/", обрабатываем как обычную дробь
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

	// Если содержит точку, считаем как десятичную дробь
	if strings.Contains(s, ".") {
		floatVal, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, errors.New("не удалось разобрать десятичное число")
		}

		// Преобразуем десятичную дробь в обыкновенную
		// Например: 0.25 → 25 / 100 → 1 / 4
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

	// Иначе — просто целое число
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return NewFractionNumber(n, 1)
}

func (f *FractionNumber) Denominator() int64 {
	return f.denominator
}

func (f *FractionNumber) Numerator() int64 {
	return f.numerator
}

func (f *FractionNumber) Add(other *FractionNumber) *FractionNumber {
	m := lcm(f.denominator, other.denominator)
	sum := &FractionNumber{
		numerator:   f.numerator*(m/f.denominator) + other.numerator*(m/other.denominator),
		denominator: m,
	}
	return sum.shrink()
}

func (f *FractionNumber) Sub(other *FractionNumber) *FractionNumber {
	other.numerator *= -1
	return f.Add(other).shrink()
}

func (f *FractionNumber) Mul(other *FractionNumber) *FractionNumber {
	frac, _ := NewFractionNumber(f.numerator*other.numerator, f.denominator*other.denominator)
	return frac.shrink()
}

func (f *FractionNumber) Div(other *FractionNumber) (*FractionNumber, error) {
	frac, err := NewFractionNumber(f.numerator*other.denominator, f.denominator*other.numerator)
	if err != nil {
		err = ErrDivideByZero
	}
	return frac.shrink(), err
}

func (f *FractionNumber) Square() *FractionNumber {
	return f.Mul(f)
}

func (f *FractionNumber) Reverse() *FractionNumber {
	return f.Mul(RevOneValue).shrink()
}

func (f *FractionNumber) Inverse() (*FractionNumber, error) {
	inv, err := OneValue.Div(f)
	if err != nil {
		return nil, err
	}
	return inv.shrink(), nil
}

func (f *FractionNumber) Equal(other *FractionNumber) bool {
	return f.numerator == other.numerator && f.denominator == other.denominator
}

func (f *FractionNumber) NotEqual(other *FractionNumber) bool {
	return !f.Equal(other)
}

func (f *FractionNumber) Float64() float64 {
	f.shrink()
	return float64(f.numerator) / float64(f.denominator)
}

func (f *FractionNumber) String() string {
	if f.denominator == 1 {
		return fmt.Sprintf("%d", f.numerator)
	}
	f.shrink()
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}

func (f *FractionNumber) shrink() *FractionNumber {
	gcf := gcd(abs(f.numerator), f.denominator)
	return &FractionNumber{f.numerator / gcf, f.denominator / gcf}
}

func (f *FractionNumber) LessThan(other *FractionNumber) bool {
	return (f.numerator * other.denominator) < (other.Numerator() * f.denominator)
}

func (f *FractionNumber) GreaterThan(other *FractionNumber) bool {
	return (f.numerator * other.denominator) > (other.Numerator() * f.denominator)
}

func (f *FractionNumber) Abs() *FractionNumber {
	if f.numerator < 0 {
		return &FractionNumber{
			numerator:   -f.numerator,
			denominator: f.denominator,
		}
	}
	return f.shrink()
}

func abs[T integer](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(n1, n2 int64) int64 {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}

func lcm(n1, n2 int64) int64 {
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	return n1 * (n2 / gcd(n1, n2))
}
