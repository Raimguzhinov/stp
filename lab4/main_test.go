package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMatrix(t *testing.T) {
	type args struct {
		data [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMatrix(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Add(t *testing.T) {
	tests := []struct {
		name     string
		matrix1  [][]int
		matrix2  [][]int
		expected [][]int
	}{
		{
			name:     "Normal addition",
			matrix1:  [][]int{{1, 2}, {3, 4}},
			matrix2:  [][]int{{5, 6}, {7, 8}},
			expected: [][]int{{6, 8}, {10, 12}},
		},
		{
			name:     "Zero matrix addition",
			matrix1:  [][]int{{0, 0}, {0, 0}},
			matrix2:  [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{1, 2}, {3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix1, _ := NewMatrix(tt.matrix1)
			matrix2, _ := NewMatrix(tt.matrix2)

			resultMatrix, err := matrix1.Add(matrix2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, resultMatrix.data)
		})
	}
}

func TestMatrix_Sub(t *testing.T) {
	tests := []struct {
		name     string
		matrix1  [][]int
		matrix2  [][]int
		expected [][]int
	}{
		{
			name:     "Normal subtraction",
			matrix1:  [][]int{{5, 6}, {7, 8}},
			matrix2:  [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{4, 4}, {4, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix1, _ := NewMatrix(tt.matrix1)
			matrix2, _ := NewMatrix(tt.matrix2)

			resultMatrix, err := matrix1.Sub(matrix2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, resultMatrix.data)
		})
	}
}

func TestMatrix_Mul(t *testing.T) {
	tests := []struct {
		name     string
		matrix1  [][]int
		matrix2  [][]int
		expected [][]int
	}{
		{
			name:     "Normal multiplication",
			matrix1:  [][]int{{1, 2}, {3, 4}},
			matrix2:  [][]int{{5, 6}, {7, 8}},
			expected: [][]int{{19, 22}, {43, 50}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix1, _ := NewMatrix(tt.matrix1)
			matrix2, _ := NewMatrix(tt.matrix2)

			resultMatrix, err := matrix1.Mul(matrix2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, resultMatrix.data)
		})
	}
}

func TestMatrix_Transpose(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name:     "Normal transpose",
			matrix:   [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{1, 3}, {2, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, _ := NewMatrix(tt.matrix)

			resultMatrix, err := matrix.Transpose()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, resultMatrix.data)
		})
	}
}

func TestMatrix_Min(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name:     "Matrix with negative values",
			matrix:   [][]int{{1, 2}, {3, -4}},
			expected: -4,
		},
		{
			name:     "All positive values",
			matrix:   [][]int{{1, 2}, {3, 4}},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, _ := NewMatrix(tt.matrix)

			resultMin := matrix.Min()
			assert.Equal(t, tt.expected, resultMin)
		})
	}
}

func TestMatrix_ToString(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected string
	}{
		{
			name:     "Normal case",
			matrix:   [][]int{{1, 2}, {3, 4}},
			expected: "{{[1 2]}, {[3 4]}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, _ := NewMatrix(tt.matrix)

			resultString := matrix.ToString()
			assert.Equal(t, tt.expected, resultString)
		})
	}
}

func TestMatrix_Get(t *testing.T) {
	tests := []struct {
		name          string
		matrix        [][]int
		i, j          int
		expectedValue int
		expectError   bool
	}{
		{
			name:          "Normal case",
			matrix:        [][]int{{1, 2}, {3, 4}},
			i:             1,
			j:             1,
			expectedValue: 4,
			expectError:   false,
		},
		{
			name:        "Out of bounds",
			matrix:      [][]int{{1, 2}, {3, 4}},
			i:           2,
			j:           2,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, _ := NewMatrix(tt.matrix)

			resultValue, err := matrix.Get(tt.i, tt.j)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedValue, resultValue)
			}
		})
	}
}

func TestMatrix_Equals(t *testing.T) {
	tests := []struct {
		name     string
		matrix1  [][]int
		matrix2  [][]int
		expected bool
	}{
		{
			name:     "Equal matrices",
			matrix1:  [][]int{{1, 2}, {3, 4}},
			matrix2:  [][]int{{1, 2}, {3, 4}},
			expected: true,
		},
		{
			name:     "Unequal matrices",
			matrix1:  [][]int{{1, 2}, {3, 4}},
			matrix2:  [][]int{{5, 6}, {7, 8}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix1, _ := NewMatrix(tt.matrix1)
			matrix2, _ := NewMatrix(tt.matrix2)

			assert.Equal(t, tt.expected, matrix1.Equals(matrix2))
		})
	}
}

func TestMatrix_RowsAndCols(t *testing.T) {
	tests := []struct {
		name         string
		matrix       [][]int
		expectedRows int
		expectedCols int
	}{
		{
			name:         "Normal case",
			matrix:       [][]int{{1, 2}, {3, 4}},
			expectedRows: 2,
			expectedCols: 2,
		},
		{
			name:         "Single row matrix",
			matrix:       [][]int{{1, 2}},
			expectedRows: 1,
			expectedCols: 2,
		},
		{
			name:         "Single column matrix",
			matrix:       [][]int{{1}, {2}},
			expectedRows: 2,
			expectedCols: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, _ := NewMatrix(tt.matrix)

			assert.Equal(t, tt.expectedRows, matrix.Rows(), "Number of rows is incorrect")
			assert.Equal(t, tt.expectedCols, matrix.Cols(), "Number of columns is incorrect")
		})
	}
}
