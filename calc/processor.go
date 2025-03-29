package main

import "errors"

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Execute(op string, a, b *FractionNumber) (*FractionNumber, error) {
	switch op {
	case LabelPlus:
		return a.Add(b), nil
	case LabelMinus:
		return a.Sub(b), nil
	case LabelMultiply:
		return a.Mul(b), nil
	case LabelDivide:
		return a.Div(b)
	default:
		return nil, errors.New("неизвестная операция")
	}
}

func (p *Processor) ApplyUnary(op string, a *FractionNumber) (*FractionNumber, error) {
	switch op {
	case LabelSqr:
		return a.Square(), nil
	case LabelInverse:
		return a.Inverse()
	default:
		return nil, errors.New("неизвестная унарная операция")
	}
}
