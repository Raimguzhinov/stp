package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatrixOperations_MinOfThree(t *testing.T) {
	m := MatrixOperations{}
	tests := []struct {
		name     string
		a, b, c  float64
		expected float64
	}{
		{
			name:     "All distinct numbers",
			a:        3.0,
			b:        -5.0,
			c:        1.0,
			expected: -5.0,
		},
		{
			name:     "All zeros",
			a:        0.0,
			b:        0.0,
			c:        0.0,
			expected: 0.0,
		},
		{
			name:     "First and second are equal",
			a:        5.0,
			b:        5.0,
			c:        10.0,
			expected: 5.0,
		},
		{
			name:     "First and third are equal",
			a:        7.0,
			b:        8.5,
			c:        7.5,
			expected: 7.0,
		},
		{
			name:     "Second and third are equal",
			a:        9.0,
			b:        4.0,
			c:        4.0,
			expected: 4.0,
		},
		{
			name:     "First is minimum",
			a:        1.1,
			b:        2.2,
			c:        3.3,
			expected: 1.1,
		},
		{
			name:     "Second is minimum",
			a:        10.0,
			b:        -10.0,
			c:        5.0,
			expected: -10.0,
		},
		{
			name:     "Third is minimum",
			a:        10.0,
			b:        15.0,
			c:        -5.0,
			expected: -5.0,
		},
		{
			name:     "Two minimums (a == b)",
			a:        5.0,
			b:        5.0,
			c:        10.0,
			expected: 5.0,
		},
		{
			name:     "Two minimums (b == c)",
			a:        10.0,
			b:        5.0,
			c:        5.0,
			expected: 5.0,
		},
		{
			name:     "Two minimums (a == c)",
			a:        5.0,
			b:        10.0,
			c:        5.0,
			expected: 5.0,
		},
		{
			name:     "All equal",
			a:        1.0,
			b:        1.0,
			c:        1.0,
			expected: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := m.MinOfThree(tt.a, tt.b, tt.c)
			assert.Equal(t, tt.expected, result, "Expected minimum did not match")
		})
	}
}

func TestMatrixOperations_SumEvenIndexElements(t *testing.T) {
	m := MatrixOperations{}
	tests := []struct {
		name        string
		input       [][]float64
		expectedSum float64
		expectError bool
	}{
		{
			name:        "Normal matrix",
			input:       [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}},
			expectedSum: 25.0,
			expectError: false,
		},
		{
			name:        "Empty matrix",
			input:       [][]float64{},
			expectedSum: 0.0,
			expectError: true,
		},
		{
			name:        "Single element matrix",
			input:       [][]float64{{42.0}},
			expectedSum: 42.0,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, err := m.SumEvenIndexElements(tt.input)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedSum, sum, "Expected sum did not match")
			}
		})
	}
}

func TestMatrixOperations_MaxOnAndBelowDiagonal(t *testing.T) {
	m := MatrixOperations{}
	tests := []struct {
		name        string
		input       [][]float64
		expectedMax float64
		expectError bool
	}{
		{
			name:        "Normal matrix",
			input:       [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}},
			expectedMax: 9.0,
			expectError: false,
		},
		{
			name:        "Empty matrix",
			input:       [][]float64{},
			expectedMax: 0.0,
			expectError: true,
		},
		{
			name:        "Single element matrix",
			input:       [][]float64{{42.0}},
			expectedMax: 42.0,
			expectError: false,
		},
		{
			name:        "Rectangular matrix",
			input:       [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}},
			expectedMax: 5.0,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxVal, err := m.MaxOnAndBelowDiagonal(tt.input)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedMax, maxVal, "Expected max value did not match")
			}
		})
	}
}
