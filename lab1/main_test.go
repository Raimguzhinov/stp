package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_productOfNonZero(t *testing.T) {
	tests := []struct {
		name     string
		v        []int
		ind      []int
		expected int
		errMsg   string
	}{
		{
			name:     "Several non-zero values",
			v:        []int{2, 0, 5, 7, 0},
			ind:      []int{0, 2, 3},
			expected: 70,
			errMsg:   "",
		},
		{
			name:     "All values zero",
			v:        []int{0, 0, 0},
			ind:      []int{0, 1, 2},
			expected: 0,
			errMsg:   "",
		},
		{
			name:     "Empty arrays",
			v:        []int{},
			ind:      []int{},
			expected: 0,
			errMsg:   "empty array",
		},
		{
			name:     "Nil arrays",
			v:        nil,
			ind:      nil,
			expected: 0,
			errMsg:   "empty array",
		},
		{
			name:     "Empty V array",
			v:        []int{1, 2, 3},
			ind:      []int{},
			expected: 0,
			errMsg:   "empty array",
		},
		{
			name:     "Nil V array",
			v:        []int{1, 2, 3},
			ind:      nil,
			expected: 0,
			errMsg:   "empty array",
		},
		{
			name:     "Empty Ind array",
			v:        []int{},
			ind:      []int{1, 2, 3},
			expected: 0,
			errMsg:   "empty array",
		},
		{
			name:     "Nil Ind array",
			v:        nil,
			ind:      []int{1, 2, 3},
			expected: 0,
			errMsg:   "empty array",
		},
		{
			name:     "One non-zero value",
			v:        []int{0, 3, 0},
			ind:      []int{1},
			expected: 3,
			errMsg:   "",
		},
		{
			name:     "Out of up bound index",
			v:        []int{1, 2, 3},
			ind:      []int{0, 3},
			expected: 0,
			errMsg:   "index 3 is out of bounds",
		},
		{
			name:     "Out of down bound index",
			v:        []int{1, 2, 3},
			ind:      []int{0, -1},
			expected: 0,
			errMsg:   "index -1 is out of bounds",
		},
		{
			name:     "Mixed zero and non-zero values",
			v:        []int{1, 0, 2, 3},
			ind:      []int{0, 1, 2, 3},
			expected: 6,
			errMsg:   "",
		},
		{
			name:     "No non-zero values",
			v:        []int{0, 0, 0, 0},
			ind:      []int{0, 1, 2, 3},
			expected: 0,
			errMsg:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := productOfNonZero(tt.v, tt.ind)

			if tt.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unexpected product result")
			}
		})
	}
}

func Test_findMinAndIndex(t *testing.T) {
	tests := []struct {
		name        string
		arr         []int
		expectedVal int
		expectedIdx int
		expectError bool
	}{
		{
			name:        "Normal array",
			arr:         []int{5, 3, 8, 1, 4, 7},
			expectedVal: 1,
			expectedIdx: 3,
			expectError: false,
		},
		{
			name:        "Single element array",
			arr:         []int{42},
			expectedVal: 42,
			expectedIdx: 0,
			expectError: false,
		},
		{
			name:        "All elements are the same",
			arr:         []int{7, 7, 7, 7},
			expectedVal: 7,
			expectedIdx: 0,
			expectError: false,
		},
		{
			name:        "Empty array",
			arr:         []int{},
			expectedVal: 0,
			expectedIdx: -1,
			expectError: true,
		},
		{
			name:        "Negative values",
			arr:         []int{-5, -10, -3, -8},
			expectedVal: -10,
			expectedIdx: 1,
			expectError: false,
		},
		{
			name:        "Mixed positive and negative values",
			arr:         []int{10, -3, 7, 0, -15, 20},
			expectedVal: -15,
			expectedIdx: 4,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minVal, minIdx, err := findMinAndIndex(tt.arr)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedVal, minVal)
				assert.Equal(t, tt.expectedIdx, minIdx)
			}
		})
	}
}

func Test_reverseArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "Normal case",
			input:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: []float64{5.5, 4.4, 3.3, 2.2, 1.1},
		},
		{
			name:     "Single element",
			input:    []float64{42.42},
			expected: []float64{42.42},
		},
		{
			name:     "Empty array",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "Nil array",
			input:    nil,
			expected: []float64{},
		},
		{
			name:     "Two elements",
			input:    []float64{3.3, 4.4},
			expected: []float64{4.4, 3.3},
		},
		{
			name:     "Negative and positive values",
			input:    []float64{-1.1, 2.2, -3.3, 4.4},
			expected: []float64{4.4, -3.3, 2.2, -1.1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]float64, len(tt.input))
			copy(inputCopy, tt.input)
			reverseArray(inputCopy)
			assert.Equal(t, tt.expected, inputCopy)
		})
	}
}
