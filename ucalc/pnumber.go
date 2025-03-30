package main

import "stp/lab5/upnumber"

type TPNumberAdapter struct {
	tp *upnumber.TPNumber
}

func NewTPNumberAdapter(tp *upnumber.TPNumber) *TPNumberAdapter {
	return &TPNumberAdapter{tp: tp}
}

func (t *TPNumberAdapter) Add(n Number) Number {
	r, _ := t.tp.Add(n.(*TPNumberAdapter).tp)
	return NewTPNumberAdapter(r)
}

func (t *TPNumberAdapter) Sub(n Number) Number {
	r, _ := t.tp.Sub(n.(*TPNumberAdapter).tp)
	return NewTPNumberAdapter(r)
}

func (t *TPNumberAdapter) Mul(n Number) Number {
	r, _ := t.tp.Mul(n.(*TPNumberAdapter).tp)
	return NewTPNumberAdapter(r)
}

func (t *TPNumberAdapter) Div(n Number) (Number, error) {
	r, err := t.tp.Div(n.(*TPNumberAdapter).tp)
	if err != nil {
		return nil, err
	}
	return NewTPNumberAdapter(r), nil
}

func (t *TPNumberAdapter) Square() Number {
	r, _ := t.tp.Square()
	return NewTPNumberAdapter(r)
}

func (t *TPNumberAdapter) Inverse() (Number, error) {
	r, err := t.tp.Inverse()
	if err != nil {
		return nil, err
	}
	return NewTPNumberAdapter(r), nil
}

func (t *TPNumberAdapter) String() string {
	return t.tp.ToString()
}

func (t *TPNumberAdapter) Float64() float64 {
	return t.tp.GetValue()
}

func (t *TPNumberAdapter) Copy() Number {
	return NewTPNumberAdapter(t.tp.Copy())
}
