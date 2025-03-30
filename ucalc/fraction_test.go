package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewFractionNumber(t *testing.T) {
	f, err := NewFractionNumber(4, 8)
	require.NoError(t, err)
	require.Equal(t, int64(1), f.Numerator())
	require.Equal(t, int64(2), f.Denominator())

	_, err = NewFractionNumber(1, 0)
	require.ErrorIs(t, err, ErrZeroDenominator)
}

func TestFractionComplexSequence(t *testing.T) {
	a, err := NewFractionNumber(3, 4) // 3/4
	require.NoError(t, err)

	b, err := NewFractionNumber(5, 6) // 5/6
	require.NoError(t, err)

	// (3/4 + 5/6) = 19/12
	sum := a.Add(b)
	require.Equal(t, "19/12", sum.String())

	// (19/12 - 2/3) = 11/12
	c, _ := NewFractionNumber(2, 3)
	diff := sum.Sub(c)
	require.Equal(t, "11/12", diff.String())

	// (11/12 * 4/11) = 44/132 = 1/3
	d, _ := NewFractionNumber(4, 11)
	prod := diff.Mul(d)
	require.Equal(t, "1/3", prod.String())

	// (1/3 / 2/3) = 1/3 * 3/2 = 3/6 = 1/2
	e, _ := NewFractionNumber(2, 3)
	div, err := prod.Div(e)
	require.NoError(t, err)
	require.Equal(t, "1/2", div.String())

	// Inverse(1/2) = 2
	inv, err := div.Inverse()
	require.NoError(t, err)
	require.Equal(t, "2", inv.String())

	// Square(2) = 4
	sqr := inv.Square()
	require.Equal(t, "4", sqr.String())
}

func TestFraction_Add(t *testing.T) {
	a, _ := NewFractionNumber(1, 2)
	b, _ := NewFractionNumber(1, 3)
	c := a.Add(b).(*FractionNumber)
	require.Equal(t, "5/6", c.String())
}

func TestFraction_Sub(t *testing.T) {
	a, _ := NewFractionNumber(3, 4)
	b, _ := NewFractionNumber(1, 4)
	c := a.Sub(b).(*FractionNumber)
	require.Equal(t, "1/2", c.String())
}

func TestFraction_Mul(t *testing.T) {
	a, _ := NewFractionNumber(2, 3)
	b, _ := NewFractionNumber(3, 4)
	c := a.Mul(b).(*FractionNumber)
	require.Equal(t, "1/2", c.String())
}

func TestFraction_Div(t *testing.T) {
	a, _ := NewFractionNumber(1, 2)
	b, _ := NewFractionNumber(1, 4)
	c, err := a.Div(b)
	require.NoError(t, err)
	require.Equal(t, "2", c.String())

	_, err = a.Div(&FractionNumber{0, 1})
	require.ErrorIs(t, err, ErrDivideByZero)
}

func TestFraction_Square(t *testing.T) {
	a, _ := NewFractionNumber(2, 3)
	sq := a.Square().(*FractionNumber)
	require.Equal(t, "4/9", sq.String())
}

func TestFraction_Inverse(t *testing.T) {
	a, _ := NewFractionNumber(2, 3)
	inv, err := a.Inverse()
	require.NoError(t, err)
	require.Equal(t, "3/2", inv.String())
}

func TestFraction_Copy(t *testing.T) {
	a, _ := NewFractionNumber(5, 6)
	b := a.Copy().(*FractionNumber)
	require.Equal(t, "5/6", b.String())
	b.numerator = 2
	require.Equal(t, int64(5), a.numerator, "Original should not be modified")
}

func TestFraction_Float64(t *testing.T) {
	a, _ := NewFractionNumber(1, 2)
	require.InDelta(t, 0.5, a.Float64(), 1e-9)
}

func TestFraction_shrink(t *testing.T) {
	f := &FractionNumber{10, 20}
	shrunk := f.shrink()
	require.Equal(t, int64(1), shrunk.Numerator())
	require.Equal(t, int64(2), shrunk.Denominator())
}

func TestNegativeNormalization(t *testing.T) {
	f, err := NewFractionNumber(-4, -8)
	require.NoError(t, err)
	require.Equal(t, "1/2", f.String())

	f2, err := NewFractionNumber(4, -8)
	require.NoError(t, err)
	require.Equal(t, "-1/2", f2.String())
}

func TestZeroNumerator(t *testing.T) {
	f, err := NewFractionNumber(0, 5)
	require.NoError(t, err)
	require.Equal(t, "0", f.String())
}
