package uproc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessor(t *testing.T) {
	defaultValue := 0.9
	proc := NewTProc(defaultValue)

	// Тест Reset
	proc.Reset()
	assert.Equal(t, 0.0, proc.GetLOpAndRes())
	assert.Equal(t, 0.0, proc.GetROp())
	assert.Equal(t, None, proc.GetOperation())

	// Тест LOpAndResSet и GetLOpAndRes
	proc.LOpAndResSet(10.4)
	assert.InDelta(t, 10.4, proc.GetLOpAndRes(), 1e-9)

	// Тест ROpSet и GetROp
	proc.ROpSet(5.2)
	assert.InDelta(t, 5.2, proc.GetROp(), 1e-9)

	// Тест OperationSet и GetOperation
	proc.OperationSet(Add)
	assert.Equal(t, Add, proc.GetOperation())

	// Тест OperationRun с Add
	assert.NoError(t, proc.OperationRun())
	assert.InDelta(t, 15.6, proc.GetLOpAndRes(), 1e-9)

	// Тест OperationRun с Sub
	proc.OperationSet(Sub)
	proc.ROpSet(3)
	assert.NoError(t, proc.OperationRun())
	assert.InDelta(t, 12.6, proc.GetLOpAndRes(), 1e-9)

	// Тест OperationRun с Mul
	proc.OperationSet(Mul)
	proc.ROpSet(2)
	assert.NoError(t, proc.OperationRun())
	assert.InDelta(t, 25.2, proc.GetLOpAndRes(), 1e-9)

	// Тест OperationRun с Dvd
	proc.OperationSet(Dvd)
	proc.ROpSet(6)
	assert.NoError(t, proc.OperationRun())
	assert.InDelta(t, 4.2, proc.GetLOpAndRes(), 1e-9)

	// Ошибка деления на ноль
	proc.ROpSet(0)
	assert.EqualError(t, proc.OperationRun(), "division by zero")

	// Тест FuncRun с Rev
	proc.ROpSet(4)
	assert.NoError(t, proc.FuncRun(Rev))
	assert.InDelta(t, 0.25, proc.GetROp(), 1e-9)

	// Тест FuncRun с Sqr
	assert.NoError(t, proc.FuncRun(Sqr))
	assert.InDelta(t, 0.0625, proc.GetROp(), 1e-9)

	// Ошибка FuncRun с некорректной функцией
	assert.EqualError(t, proc.FuncRun(TFunc(42)), "invalid function")

	// Ошибка OperationRun с некорректной операцией
	proc.OperationSet(TOperation(42))
	assert.EqualError(t, proc.OperationRun(), "invalid operation")
}
