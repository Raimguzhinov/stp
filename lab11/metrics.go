package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Metrics представляет результаты вычислений.
type Metrics struct {
	Length              float64
	Variance            float64
	StdDeviation        float64
	RelativeError       float64
	TheoreticalLength   float64
	TheoreticalVariance float64
	TheoreticalStdDev   float64
	TheoreticalRelError float64
}

// calculateTheoreticalValues рассчитывает теоретические значения.
func calculateTheoreticalValues(operators, operands int) Metrics {
	eta := operators + operands
	length := 0.9 * float64(eta) * math.Log2(float64(eta))
	variance := (math.Pi * math.Pi * float64(eta*eta)) / 6
	stdDev := math.Sqrt(variance)
	relError := 1 / (2 * math.Log2(float64(eta)))

	return Metrics{
		TheoreticalLength:   length,
		TheoreticalVariance: variance,
		TheoreticalStdDev:   stdDev,
		TheoreticalRelError: relError,
	}
}

// simulateProgramGeneration моделирует процесс написания программы.
func simulateProgramGeneration(eta int) float64 {
	seen := make(map[int]struct{})
	rand.Seed(time.Now().UnixNano())

	totalDraws := 0
	for len(seen) < eta {
		draw := rand.Intn(eta)
		seen[draw] = struct{}{}
		totalDraws++
	}

	return float64(totalDraws)
}

// compareMetrics сравнивает теоретические и эмпирические значения.
func compareMetrics(operators, operands int) {
	eta := operators + operands
	theoretical := calculateTheoreticalValues(operators, operands)
	simulatedLength := simulateProgramGeneration(eta)

	fmt.Printf("=== Metrics for Dictionary Size (η) = %d ===\n", eta)
	fmt.Printf("Theoretical Length: %.2f\n", theoretical.TheoreticalLength)
	fmt.Printf("Simulated Length: %.2f\n", simulatedLength)
	fmt.Printf("Theoretical Variance: %.2f\n", theoretical.TheoreticalVariance)
	fmt.Printf("Theoretical Std. Dev.: %.2f\n", theoretical.TheoreticalStdDev)
	fmt.Printf("Theoretical Rel. Error: %.2f%%\n", theoretical.TheoreticalRelError*100)
	fmt.Printf("--------------------------\n")
}
