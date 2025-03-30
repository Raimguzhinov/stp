package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEditor_InputAndBasicOps(t *testing.T) {
	e := NewEditor()
	require.Equal(t, "1", e.Input("1"))
	require.Equal(t, "12", e.Input("2"))
	require.Equal(t, "12/", e.Input(LabelFracSep))
	require.Equal(t, "12/3", e.Input("3"))
	require.Equal(t, "", e.Input(LabelPlus)) // Операция очищает буфер
	require.NotNil(t, e.operand1)
	require.Equal(t, "+", e.operation)
}

func TestEditor_NegateAndClear(t *testing.T) {
	e := NewEditor()
	require.Equal(t, "5", e.Input("5"))
	require.Equal(t, "-5", e.Input(LabelNegate))
	require.Equal(t, "5", e.Input(LabelNegate)) // повторно — убирает минус
	require.Equal(t, "", e.Input(LabelClear))
	require.Nil(t, e.operand1)
	require.Empty(t, e.buffer)
}

func TestEditor_Backspace(t *testing.T) {
	e := NewEditor()
	e.Input("7")
	e.Input("8")
	require.Equal(t, "78", e.buffer)
	e.Input(LabelBack)
	require.Equal(t, "7", e.buffer)
	e.Input(LabelBack)
	require.Equal(t, "", e.buffer)
	e.Input(LabelBack)
	require.Equal(t, "", e.buffer) // безопасный вызов на пустом буфере
}

func TestEditor_Input_Default(t *testing.T) {
	e := NewEditor()
	require.Equal(t, "3", e.Input("3"))
	require.Equal(t, "30", e.Input("0"))
}

func TestEditor_SetOrReplaceChar(t *testing.T) {
	e := NewEditor()
	e.buffer = "1."
	e.setOrReplaceChar(".", "0")
	require.Equal(t, "1.", e.buffer)

	e.buffer = ""
	e.setOrReplaceChar(".", "0")
	require.Equal(t, "0.", e.buffer)
}

func TestEditor_OperatorAfterLastResult(t *testing.T) {
	e := NewEditor()
	e.lastResult, _ = ParseNumber("1/2")

	require.Equal(t, "", e.Input(LabelPlus)) // buffer пустой, lastResult задан
	require.Equal(t, "+", e.operation)
	require.NotNil(t, e.operand1)
}

func TestEditor_InvalidFractionParsing(t *testing.T) {
	e := NewEditor()
	e.Input("bad")
	result := e.Input(LabelPlus) // Попытка спарсить "bad"
	require.Equal(t, "Ошибка", result)
}

func TestEditor_InputDotAndFrac(t *testing.T) {
	e := NewEditor()
	// . с пустым буфером
	require.Equal(t, "0.", e.Input(LabelDot))
	// Повторная точка — игнорируется
	require.Equal(t, "0.", e.Input(LabelDot))
	// / после числа
	require.Equal(t, "0.1", e.Input("1"))
	require.Equal(t, "0.1", e.Input(LabelFracSep)) // повторная дробь — игнорируется
	e.clear()
	// / с пустым буфером
	require.Equal(t, "1/", e.Input(LabelFracSep))
	// Повторная дробь — игнорируется
	require.Equal(t, "1/", e.Input(LabelFracSep))
	// . после числа
	require.Equal(t, "1/2", e.Input("2"))
	require.Equal(t, "1/2", e.Input(LabelDot)) // точка — игнорируется
}
