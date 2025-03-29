package main

import (
	"errors"
	"fmt"
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

	ZeroValue   = &FractionNumber{0, 1}
	OneValue    = &FractionNumber{1, 1}
	RevOneValue = &FractionNumber{-1, 1}
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
	return &FractionNumber{n / gcf, d / gcf}, nil
}

func (f *FractionNumber) Denominator() int64 {
	return f.denominator
}

func (f *FractionNumber) Numerator() int64 {
	return f.numerator
}

func (f *FractionNumber) Add(other Number) Number {
	o := other.(*FractionNumber)
	m := lcm(f.denominator, o.denominator)
	sum := &FractionNumber{
		numerator:   f.numerator*(m/f.denominator) + o.numerator*(m/o.denominator),
		denominator: m,
	}
	return sum.shrink()
}

func (f *FractionNumber) Sub(other Number) Number {
	o := other.(*FractionNumber)
	sub := f.Add(&FractionNumber{-o.numerator, o.denominator})
	sub.(*FractionNumber).shrink()
	return sub
}

func (f *FractionNumber) Mul(other Number) Number {
	o := other.(*FractionNumber)
	frac, _ := NewFractionNumber(f.numerator*o.numerator, f.denominator*o.denominator)
	return frac.shrink()
}

func (f *FractionNumber) Div(other Number) (Number, error) {
	o := other.(*FractionNumber)
	frac, err := NewFractionNumber(f.numerator*o.denominator, f.denominator*o.numerator)
	if err != nil {
		return nil, ErrDivideByZero
	}
	return frac.shrink(), nil
}

func (f *FractionNumber) Square() Number {
	return f.Mul(f)
}

func (f *FractionNumber) Inverse() (Number, error) {
	return OneValue.Div(f)
}

func (f *FractionNumber) Copy() Number {
	fcopy := &FractionNumber{numerator: f.Numerator(), denominator: f.Denominator()}
	return fcopy.shrink()
}

func (f *FractionNumber) Float64() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func (f *FractionNumber) String() string {
	if f.denominator == 1 {
		return fmt.Sprintf("%d", f.numerator)
	}
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}

func (f *FractionNumber) shrink() *FractionNumber {
	gcf := gcd(abs(f.numerator), f.denominator)
	return &FractionNumber{f.numerator / gcf, f.denominator / gcf}
}

func abs[T integer](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}
