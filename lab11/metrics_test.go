package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCalculateTheoreticalValues(t *testing.T) {
	tests := []struct {
		name           string
		operators      int
		operands       int
		expectedLength float64
	}{
		{"Small Dictionary", 4, 4, 0.9 * 8 * math.Log2(8)},
		{"Medium Dictionary", 16, 16, 0.9 * 32 * math.Log2(32)},
		{"Large Dictionary", 64, 64, 0.9 * 128 * math.Log2(128)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metrics := calculateTheoreticalValues(tt.operators, tt.operands)
			assert.InEpsilon(t, tt.expectedLength, metrics.TheoreticalLength, 0.01)
			assert.InEpsilon(t, (
				math.Pi*math.Pi*float64((tt.operators+tt.operands)*(tt.operators+tt.operands)))/6, metrics.TheoreticalVariance,
				0.01)
			assert.InEpsilon(t, math.Sqrt(metrics.TheoreticalVariance), metrics.TheoreticalStdDev, 0.01)
			assert.InEpsilon(t, 1/(2*math.Log2(float64(tt.operators+tt.operands))), metrics.TheoreticalRelError, 0.01)
		})
	}
}

func TestSimulateProgramGeneration(t *testing.T) {
	tests := []struct {
		name string
		eta  int
	}{
		{"Small Dictionary", 8},
		{"Medium Dictionary", 32},
		{"Large Dictionary", 128},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length := simulateProgramGeneration(tt.eta)
			assert.GreaterOrEqual(t, int(length), tt.eta)
		})
	}

	t.Run("Edge Case: Minimal Dictionary", func(t *testing.T) {
		length := simulateProgramGeneration(1)
		assert.Equal(t, float64(1), length)
	})

	t.Run("Edge Case: Invalid Dictionary Size", func(t *testing.T) {
		length := simulateProgramGeneration(0)
		assert.Equal(t, float64(0), length)
	})
}

func TestCompareMetrics(t *testing.T) {
	tests := []struct {
		name      string
		operators int
		operands  int
	}{
		{"Small Dictionary", 8, 8},
		{"Medium Dictionary", 16, 16},
		{"Large Dictionary", 32, 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eta := tt.operators + tt.operands
			theoretical := calculateTheoreticalValues(tt.operators, tt.operands)
			simulated := simulateProgramGeneration(eta)

			assert.GreaterOrEqual(t, simulated, float64(eta))
			assert.Greater(t, theoretical.TheoreticalLength, 0.0)
			assert.Greater(t, theoretical.TheoreticalVariance, 0.0)
			assert.Greater(t, theoretical.TheoreticalStdDev, 0.0)
			assert.Greater(t, theoretical.TheoreticalRelError, 0.0)
		})
	}
}
