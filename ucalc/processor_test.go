package main

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessor_Execute(t *testing.T) {
	p := NewProcessor()
	a, _ := NewFractionNumber(3, 4)
	b, _ := NewFractionNumber(1, 4)

	res, err := p.Execute(LabelPlus, a, b)
	require.NoError(t, err)
	require.Equal(t, "1", res.String())

	res, err = p.Execute(LabelMinus, a, b)
	require.NoError(t, err)
	require.Equal(t, "1/2", res.String())

	res, err = p.Execute(LabelMultiply, a, b)
	require.NoError(t, err)
	require.Equal(t, "3/16", res.String())

	res, err = p.Execute(LabelDivide, a, b)
	require.NoError(t, err)
	require.Equal(t, "3", res.String())

	_, err = p.Execute("?", a, b)
	require.Error(t, err)
	require.Equal(t, errors.New("неизвестная операция"), err)
}

func TestProcessor_ApplyUnary(t *testing.T) {
	p := NewProcessor()
	a, _ := NewFractionNumber(2, 3)

	res, err := p.ApplyUnary(LabelSqr, a)
	require.NoError(t, err)
	require.Equal(t, "4/9", res.String())

	res, err = p.ApplyUnary(LabelInverse, a)
	require.NoError(t, err)
	require.Equal(t, "3/2", res.String())

	_, err = p.ApplyUnary("?", a)
	require.Error(t, err)
	require.Equal(t, errors.New("неизвестная унарная операция"), err)
}
