package main

import (
	"math"
)

// Metrics представляет расчетные метрические характеристики ПО.
type Metrics struct {
	ModulesLowerLevel int     // Число модулей нижнего уровня (k)
	HierarchyLevels   int     // Число уровней иерархии (i)
	TotalModules      int     // Общее число модулей (K)
	ProgramLength     float64 // Длина программы (N)
	ProgramVolume     float64 // Объем программы (V)
	AssemblerCommands float64 // Длина программы в командах ассемблера (P)
	ProgrammingTime   float64 // Календарное время программирования (Tk)
	InitialErrors     float64 // Начальное количество ошибок (B)
	Reliability       float64 // Начальная надежность ПС (tn)
}

// calculateMetrics вычисляет все метрические характеристики для заданного η*₂.
func calculateMetrics(etaStar2 int, programmers int, v int, tauFactor float64) Metrics {
	// Число модулей нижнего уровня
	k := etaStar2 / 8

	// Число уровней иерархии
	i := 1
	for math.Pow(2, float64(i)) < float64(k) {
		i++
	}

	// Общее число модулей
	totalModules := etaStar2 + int(math.Pow(2, float64(i))-1)

	// Длина программы
	N := float64(totalModules) * math.Log2(float64(totalModules))

	// Объем программы
	V := float64(totalModules) * math.Pow(math.Log2(float64(totalModules)), 2)

	// Длина программы в командах ассемблера
	P := 3 * N / 8

	// Календарное время программирования
	Tk := P / (float64(v) * float64(programmers))

	// Начальное количество ошибок
	B := V / 3000

	// Надежность ПС
	tau := tauFactor * Tk
	tn := tau / (1 + math.Log(1+tau))

	return Metrics{
		ModulesLowerLevel: k,
		HierarchyLevels:   i,
		TotalModules:      totalModules,
		ProgramLength:     N,
		ProgramVolume:     V,
		AssemblerCommands: P,
		ProgrammingTime:   Tk,
		InitialErrors:     B,
		Reliability:       tn,
	}
}

// processEtaValues обрабатывает несколько значений η*₂ и возвращает метрики.
func processEtaValues(etaValues []int, programmers int, v int, tauFactor float64) []Metrics {
	var results []Metrics
	for _, etaStar2 := range etaValues {
		metrics := calculateMetrics(etaStar2, programmers, v, tauFactor)
		results = append(results, metrics)
	}
	return results
}
