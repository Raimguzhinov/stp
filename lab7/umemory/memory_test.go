package umemory

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemStateString(t *testing.T) {
	tests := []struct {
		state    MemState
		expected string
	}{
		{state: Off, expected: "_Off"},
		{state: On, expected: "_On"},
		{state: MemState(-1), expected: "Unknown"},
		{state: MemState(42), expected: "Unknown"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.state.String())
	}
}

func TestStore(t *testing.T) {
	memory := NewTMemory(0)
	memory.Store(42)
	assert.Equal(t, 42, memory.FNumber)
	assert.Equal(t, On, memory.FState)
}

func TestAdd(t *testing.T) {
	memory := NewTMemory(0)

	// Попытка добавления, когда память выключена.
	err := memory.Add(10, func(a, b int) (int, error) {
		return a + b, nil
	})
	assert.EqualError(t, err, "memory is off")

	// Успешное добавление, когда память включена.
	memory.Store(5)
	err = memory.Add(10, func(a, b int) (int, error) {
		return a + b, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, 15, memory.FNumber)
	assert.Equal(t, On, memory.FState)

	// Ошибка в addFunc.
	err = memory.Add(10, func(a, b int) (int, error) {
		return 0, errors.New("test error")
	})
	assert.EqualError(t, err, "test error")
}

func TestClear(t *testing.T) {
	memory := NewTMemory(42)
	memory.Store(42)
	memory.Clear()
	assert.Equal(t, 0, memory.FNumber)
	assert.Equal(t, Off, memory.FState)
}

func TestGetNumber(t *testing.T) {
	memory := NewTMemory(0)
	memory.Store(42)
	assert.Equal(t, 42, memory.GetNumber())
}

func TestGetState(t *testing.T) {
	memory := NewTMemory(0)
	assert.Equal(t, Off, memory.GetState())

	memory.Store(42)
	assert.Equal(t, On, memory.GetState())
}
