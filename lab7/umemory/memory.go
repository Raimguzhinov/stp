package umemory

import (
	"errors"
	"golang.org/x/exp/constraints"
)

// MemState представляет состояние памяти.
type MemState int

const (
	Off MemState = iota
	On
)

func (s MemState) String() string {
	if s < 0 || int(s) >= 2 {
		return "Unknown"
	}
	return [...]string{"_Off", "_On"}[s]
}

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// TMemory представляет параметризованную структуру памяти.
type TMemory[T Number] struct {
	FNumber T        // Число, хранящееся в памяти.
	FState  MemState // Состояние памяти (включена или выключена).
}

// NewTMemory создает новую память с начальными значениями.
func NewTMemory[T Number](defaultValue T) *TMemory[T] {
	return &TMemory[T]{
		FNumber: defaultValue,
		FState:  Off,
	}
}

// Store записывает значение в память и включает ее.
func (m *TMemory[T]) Store(value T) {
	m.FNumber = value
	m.FState = On
}

// Add добавляет значение к числу, хранящемуся в памяти.
func (m *TMemory[T]) Add(value T, addFunc func(a, b T) (T, error)) error {
	if m.FState == Off {
		return errors.New("memory is off")
	}
	result, err := addFunc(m.FNumber, value)
	if err != nil {
		return err
	}
	m.FNumber = result
	m.FState = On
	return nil
}

// Clear очищает память и выключает ее.
func (m *TMemory[T]) Clear() {
	var zeroValue T // Использование zero value для типа T
	m.FNumber = zeroValue
	m.FState = Off
}

// GetNumber возвращает текущее число из памяти.
func (m *TMemory[T]) GetNumber() T {
	return m.FNumber
}

// GetState возвращает текущее состояние памяти.
func (m *TMemory[T]) GetState() MemState {
	return m.FState
}
