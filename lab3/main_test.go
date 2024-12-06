package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNumberOperations_OddDigitsReversed(t *testing.T) {
	n := NumberOperations{}

	tests := []struct {
		name     string
		a        int
		expected int
	}{
		{
			name:     "Normal case",
			a:        12345,
			expected: 531,
		},
		{
			name:     "Single digit",
			a:        7,
			expected: 7,
		},
		{
			name:     "All odd digits",
			a:        135796,
			expected: 951,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := n.OddDigitsReversed(tt.a)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNumberOperations_MaxEvenDigitPosition(t *testing.T) {
	n := NumberOperations{}

	tests := []struct {
		name        string
		a           int
		expectedPos int
		expectError bool
	}{
		{
			name:        "Normal case",
			a:           62543,
			expectedPos: 2,
		},
		{
			name:        "No even digits",
			a:           13579,
			expectedPos: -1,
			expectError: true,
		},
		{
			name:        "All even digits",
			a:           2468,
			expectedPos: 2,
		},
		{
			name:        "Single digit",
			a:           4,
			expectedPos: -1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, err := n.MaxEvenDigitPosition(tt.a)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPos, pos)
			}
		})
	}
}

func TestNumberOperations_RotateDigits(t *testing.T) {
	n := NumberOperations{}

	tests := []struct {
		name     string
		a        int
		shift    int
		expected int
	}{
		{"Shift by 2", 123456, 2, 561234},
		{"Shift by 1", 98765, 1, 59876},
		{"Shift by 0", 24680, 0, 24680},
		{"Shift equals length", 12345, 5, 12345},
		{"Shift more than length", 12345, 7, 45123},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := n.RotateDigits(tt.a, tt.shift)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNumberOperations_SumAboveAntiDiagonal(t *testing.T) {
	n := NumberOperations{}

	tests := []struct {
		name        string
		matrix      [][]float64
		expectedSum float64
		expectError bool
	}{
		{
			name: "Normal matrix",
			matrix: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedSum: 6,
			expectError: false,
		},
		{
			name: "No even numbers",
			matrix: [][]float64{
				{1, 3, 5},
				{7, 9, 11},
				{13, 15, 17},
			},
			expectedSum: 0,
			expectError: false,
		},
		{
			name:        "Empty matrix",
			matrix:      [][]float64{},
			expectedSum: 0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, err := n.SumAboveAntiDiagonal(tt.matrix)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedSum, sum)
			}
		})
	}
}
