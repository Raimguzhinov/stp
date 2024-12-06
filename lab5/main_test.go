package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTPNumber_Add(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		base      int
		precision int
		expected  string
	}{
		{"Add same base", 5.125, 3.2, 10, 3, "8.325"},
		{"Add same base (base 16)", 10.75, 1.5, 16, 2, "12.25"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num1, _ := NewTPNumber(tt.a, tt.base, tt.precision)
			num2, _ := NewTPNumber(tt.b, tt.base, tt.precision)

			result, err := num1.Add(num2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result.GetValueString())
		})
	}
}

func TestTPNumber_Mul(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		base      int
		precision int
		expected  string
	}{
		{"Multiply same base", 5.125, 3.2, 10, 3, "16.400"},
		{"Multiply same base (base 16)", 10.0, 1.5, 16, 2, "15.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num1, _ := NewTPNumber(tt.a, tt.base, tt.precision)
			num2, _ := NewTPNumber(tt.b, tt.base, tt.precision)

			result, err := num1.Mul(num2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result.GetValueString())
		})
	}
}

func TestTPNumber_Sub(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		base      int
		precision int
		expected  string
	}{
		{"Subtract same base", 5.125, 3.2, 10, 3, "1.925"},
		{"Subtract same base (base 16)", 10.75, 1.5, 16, 2, "9.25"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num1, _ := NewTPNumber(tt.a, tt.base, tt.precision)
			num2, _ := NewTPNumber(tt.b, tt.base, tt.precision)

			result, err := num1.Sub(num2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result.GetValueString())
		})
	}
}

func TestTPNumber_Div(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		base      int
		precision int
		expected  string
		expectErr bool
	}{
		{"Divide same base", 10.0, 2.0, 10, 3, "5.0", false},
		{"Divide by zero", 5.125, 0.0, 10, 3, "0", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num1, _ := NewTPNumber(tt.a, tt.base, tt.precision)
			num2, _ := NewTPNumber(tt.b, tt.base, tt.precision)

			result, err := num1.Div(num2)

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result.GetValueString())
			}
		})
	}
}

func TestTPNumber_Inverse(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		base      int
		precision int
		expected  string
		expectErr bool
	}{
		{"Inverse of positive", 5.125, 10, 3, "0.195", false},
		{"Inverse of zero", 0.0, 10, 3, "0", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num, _ := NewTPNumber(tt.a, tt.base, tt.precision)

			result, err := num.Inverse()

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result.GetValueString())
			}
		})
	}
}

func TestTPNumber_Square(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		base      int
		precision int
		expected  float64
	}{
		{"Square of positive number", 5.125, 10, 3, 26.266},
		{"Square of negative number", -3.0, 10, 3, 9.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num, _ := NewTPNumber(tt.a, tt.base, tt.precision)

			result, err := num.Square()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result.GetValue())
		})
	}
}
