package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestControlUnit_EvaluateSimple(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("1")
	ctrl.Input(LabelFracSep)
	ctrl.Input("2")
	ctrl.Input(LabelPlus)
	ctrl.Input("1")
	ctrl.Input(LabelFracSep)
	ctrl.Input("4")
	result := ctrl.Evaluate()

	require.Equal(t, "3/4", result)
	require.Equal(t, "3/4", ctrl.editor.buffer)
}

func TestControlUnit_RepeatEquals(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("2")
	ctrl.Input(LabelPlus)
	ctrl.Input("2")
	require.Equal(t, "4", ctrl.Evaluate())
	require.Equal(t, "6", ctrl.Evaluate()) // повторное нажатие '='
	require.Equal(t, "8", ctrl.Evaluate())
}

func TestControlUnit_ApplyUnary(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("2")
	ctrl.Input(LabelFracSep)
	ctrl.Input("3")

	result := ctrl.ApplyFunction(LabelSqr)
	require.Equal(t, "4/9", result)

	result = ctrl.ApplyFunction(LabelInverse)
	require.Equal(t, "9/4", result)
}

func TestControlUnit_MemoryOps(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("1")
	ctrl.Input(LabelPlus)
	ctrl.Input("2")
	require.Equal(t, "3", ctrl.Evaluate())

	ctrl.MemorySave()
	ctrl.MemoryAdd()
	require.Equal(t, "6", ctrl.MemoryRead())

	ctrl.MemoryClear()
	require.Nil(t, ctrl.memory.Read())
}

func TestControlUnit_PasteExpression(t *testing.T) {
	ctrl := NewControlUnit()
	result := ctrl.PasteExpression("1/2 + 1/4")
	require.Equal(t, "3/4", result)

	single := ctrl.PasteExpression("5/6")
	require.Equal(t, "5/6", single)

	errCase := ctrl.PasteExpression("bad input")
	require.Equal(t, "Ошибка", errCase)
}

func TestControlUnit_CopyExpression(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("1")
	ctrl.Input(LabelPlus)
	ctrl.Input("2")
	require.Equal(t, "3", ctrl.Evaluate())
	require.Equal(t, "3", ctrl.CopyExpression())
}

func TestControlUnit_HistoryAccess(t *testing.T) {
	ctrl := NewControlUnit()
	require.Empty(t, ctrl.HistoryList())

	ctrl.Input("1")
	ctrl.Input(LabelPlus)
	ctrl.Input("2")
	ctrl.Evaluate()

	h := ctrl.HistoryList()
	require.Len(t, h, 1)
	require.Contains(t, h[0], "1 + 2 = 3")

	ctrl.DeleteHistory(0)
	require.Empty(t, ctrl.HistoryList())
}

func TestControlUnit_CopyExpressionFormats(t *testing.T) {
	ctrl := NewControlUnit()

	// Пусто
	require.Equal(t, "", ctrl.CopyExpression())

	// Частичное выражение
	ctrl.Input("1")
	ctrl.Input(LabelPlus)
	ctrl.Input("2")
	require.Contains(t, ctrl.CopyExpression(), "1 + 2")

	// После вычисления
	ctrl.Evaluate()
	require.Equal(t, "3", ctrl.CopyExpression())
}

func TestControlUnit_ApplyFunction_InvalidInput(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.editor.buffer = "не число"
	result := ctrl.ApplyFunction(LabelSqr)
	require.Equal(t, "Ошибка", result)
}

func TestControlUnit_ApplyFunction_InvalidUnaryOperation(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("2")
	result := ctrl.ApplyFunction("sqrt") // Неизвестная операция
	require.Equal(t, "Ошибка", result)
}

func TestControlUnit_Evaluate_NoOpNoRepeat(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.editor.buffer = "3"
	require.Equal(t, "3", ctrl.Evaluate())
}

func TestControlUnit_Evaluate_ParseError(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("1")
	ctrl.Input(LabelPlus)
	ctrl.editor.buffer = "abc"
	result := ctrl.Evaluate()
	require.Equal(t, "abc", result)
}

func TestControlUnit_Evaluate_ExecutionError(t *testing.T) {
	ctrl := NewControlUnit()
	ctrl.Input("1")
	ctrl.Input(LabelDivide)
	ctrl.Input("0")
	result := ctrl.Evaluate()
	require.Equal(t, "Ошибка", result)
}

func TestControlUnit_MemoryRead_Empty(t *testing.T) {
	ctrl := NewControlUnit()
	require.Equal(t, "", ctrl.MemoryRead())
}

func TestControlUnit_PasteExpression_WithEquals(t *testing.T) {
	ctrl := NewControlUnit()
	result := ctrl.PasteExpression("1/2+1/2=1")
	require.Equal(t, "1", result)

	result = ctrl.PasteExpression("1/2+1/2")
	require.Equal(t, "1", result)
}

func TestControlUnit_PasteExpression_InvalidBinaryFormat(t *testing.T) {
	ctrl := NewControlUnit()
	result := ctrl.PasteExpression("1/2+")
	require.Equal(t, "Ошибка", result)

	result = ctrl.PasteExpression("1/2+1/4+1/8")
	require.Equal(t, "Ошибка", result)
}

func TestControlUnit_PasteExpression_ExecutionError(t *testing.T) {
	ctrl := NewControlUnit()
	result := ctrl.PasteExpression("1/2 ÷ 0")
	require.Equal(t, "Ошибка", result)
}
