package main

import (
	"math"
)

// Metrics хранит результаты расчетов метрических характеристик.
type Metrics struct {
	TaskName             string  // Название задачи
	UniqueParams         int     // η*₂ - число уникальных параметров (вход/выход)
	Operators            int     // η₁ - число отдельных операторов
	Operands             int     // η₂ - число отдельных операндов
	OperatorCount        int     // N₁ - общее число вхождений операторов
	OperandCount         int     // N₂ - общее число вхождений операндов
	DictionaryLength     int     // η - длина словаря реализации (η = η₁ + η₂)
	ImplementationLength int     // N - длина реализации (N = N₁ + N₂)
	PredictedLength      float64 // N^ - предсказанная длина реализации (по Холстеду)
	PotentialVolume      float64 // V* - потенциальный объем реализации
	RealVolume           float64 // V - реальный объем реализации
	ProgramLevel         float64 // L - уровень программы через V* и V (L = V* / V)
	IntellectContent     float64 // I - интеллектуальное содержание программы
	PredictedTime1       float64 // T^1
	PredictedTime2       float64 // T^2
	PredictedTime3       float64 // T^3
	AverageLevel1        float64 // λ1
	AverageLevel2        float64 // λ2
}

// calculateMetrics вычисляет метрические характеристики для реализации алгоритма.
func calculateMetrics(taskName string, operators, operands, operatorCount, operandCount int) Metrics {
	dictionaryLength := operators + operands
	implementationLength := operatorCount + operandCount
	predictedLength := float64(dictionaryLength) * math.Log2(float64(dictionaryLength))
	potentialVolume := 2 * float64(dictionaryLength) * math.Log2(float64(dictionaryLength))
	realVolume := float64(implementationLength) * math.Log2(float64(dictionaryLength))
	programLevel := potentialVolume / realVolume
	intellectContent := programLevel * realVolume
	predictedTime1 := potentialVolume / 18 // Пример скорости написания 18 символов/сек
	predictedTime2 := predictedLength / 18
	predictedTime3 := realVolume / 18
	averageLevel1 := 2 * programLevel
	averageLevel2 := potentialVolume / realVolume

	return Metrics{
		TaskName:             taskName,
		UniqueParams:         2, // Пример: количество входных и выходных параметров.
		Operators:            operators,
		Operands:             operands,
		OperatorCount:        operatorCount,
		OperandCount:         operandCount,
		DictionaryLength:     dictionaryLength,
		ImplementationLength: implementationLength,
		PredictedLength:      predictedLength,
		PotentialVolume:      potentialVolume,
		RealVolume:           realVolume,
		ProgramLevel:         programLevel,
		IntellectContent:     intellectContent,
		PredictedTime1:       predictedTime1,
		PredictedTime2:       predictedTime2,
		PredictedTime3:       predictedTime3,
		AverageLevel1:        averageLevel1,
		AverageLevel2:        averageLevel2,
	}
}

// processTask выполняет задачу и возвращает метрические характеристики через callback.
func processTask(taskName string, task func()) Metrics {
	task() // Выполняем задачу (здесь функция ничего не возвращает, выполняется как операция).
	// Заполняем метрики для текущей задачи.
	switch taskName {
	case "Find Minimum Element":
		return calculateMetrics(taskName, 5, 3, 8, 6)
	case "Bubble Sort":
		return calculateMetrics(taskName, 6, 4, 12, 8)
	case "Binary Search":
		return calculateMetrics(taskName, 8, 6, 12, 10)
	case "Find Minimum in 2D Array":
		return calculateMetrics(taskName, 6, 4, 10, 8)
	case "Reverse Array":
		return calculateMetrics(taskName, 4, 3, 8, 5)
	case "Cyclic Shift":
		return calculateMetrics(taskName, 6, 4, 12, 9)
	case "Replace Elements":
		return calculateMetrics(taskName, 5, 4, 10, 8)
	default:
		return Metrics{}
	}
}
