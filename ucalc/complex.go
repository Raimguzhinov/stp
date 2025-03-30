package main

import (
	"errors"
	"fmt"
	"math/cmplx"
)

type ComplexNumber struct {
	val complex128
}

func NewComplexNumber(r, i float64) *ComplexNumber {
	return &ComplexNumber{val: complex(r, i)}
}

func (c *ComplexNumber) Add(n Number) Number {
	return &ComplexNumber{val: c.val + n.(*ComplexNumber).val}
}

func (c *ComplexNumber) Sub(n Number) Number {
	return &ComplexNumber{val: c.val - n.(*ComplexNumber).val}
}

func (c *ComplexNumber) Mul(n Number) Number {
	return &ComplexNumber{val: c.val * n.(*ComplexNumber).val}
}

func (c *ComplexNumber) Div(n Number) (Number, error) {
	if n.(*ComplexNumber).val == 0 {
		return nil, errors.New("деление на ноль")
	}
	return &ComplexNumber{val: c.val / n.(*ComplexNumber).val}, nil
}

func (c *ComplexNumber) Square() Number {
	return &ComplexNumber{val: c.val * c.val}
}

func (c *ComplexNumber) Inverse() (Number, error) {
	if c.val == 0 {
		return nil, errors.New("нельзя обратить ноль")
	}
	return &ComplexNumber{val: 1 / c.val}, nil
}

func (c *ComplexNumber) String() string {
	r := real(c.val)
	i := imag(c.val)
	if i >= 0 {
		return fmt.Sprintf("%.2f+i%.2f", r, i)
	}
	return fmt.Sprintf("%.2f-i%.2f", r, -i)
}

func (c *ComplexNumber) Float64() float64 {
	return cmplx.Abs(c.val)
}

func (c *ComplexNumber) Copy() Number {
	return &ComplexNumber{val: c.val}
}
