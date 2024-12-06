package uproc

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type TOperation int
type TFunc int

const (
	None TOperation = iota
	Add
	Sub
	Mul
	Dvd
)

func (op TOperation) String() string {
	if op < 0 || int(op) >= 5 {
		return "Unknown"
	}
	return [...]string{"None", "+", "-", "×", "÷"}[op]
}

const (
	Rev TFunc = iota
	Sqr
)

type Number interface {
	constraints.Integer | constraints.Float
}

// TProc представляет параметризованную структуру процессора.
type TProc[T Number] struct {
	LOpAndRes T          // Левый операнд и результат.
	ROp       T          // Правый операнд.
	Operation TOperation // Текущая операция.
}

// NewTProc создает новый процессор с начальными значениями.
func NewTProc[T Number](defaultValue T) *TProc[T] {
	return &TProc[T]{
		LOpAndRes: defaultValue,
		ROp:       defaultValue,
		Operation: None,
	}
}

// Reset сбрасывает процессор в начальное состояние.
func (p *TProc[T]) Reset() {
	var zeroValue T // Использование zero value для типа T
	p.LOpAndRes = zeroValue
	p.ROp = zeroValue
	p.Operation = None
}

// OperationClear сбрасывает текущую операцию.
func (p *TProc[T]) OperationClear() {
	p.Operation = None
}

// OperationSet устанавливает текущую операцию.
func (p *TProc[T]) OperationSet(op TOperation) {
	p.Operation = op
}

// LOpAndResSet устанавливает левый операнд.
func (p *TProc[T]) LOpAndResSet(value T) {
	p.LOpAndRes = value
}

// ROpSet устанавливает правый операнд.
func (p *TProc[T]) ROpSet(value T) {
	p.ROp = value
}

// OperationRun выполняет текущую операцию.
func (p *TProc[T]) OperationRun() error {
	if p.Operation == None {
		return errors.New("no operation set")
	}
	switch p.Operation {
	case Add:
		p.LOpAndRes += p.ROp
	case Sub:
		p.LOpAndRes -= p.ROp
	case Mul:
		p.LOpAndRes *= p.ROp
	case Dvd:
		if p.ROp == 0 {
			return errors.New("division by zero")
		}
		p.LOpAndRes /= p.ROp
	default:
		return errors.New("invalid operation")
	}
	return nil
}

// FuncRun выполняет функцию над правым операндом.
func (p *TProc[T]) FuncRun(f TFunc) error {
	switch f {
	case Rev:
		if p.ROp == 0 {
			return errors.New("division by zero")
		}
		p.ROp = 1 / p.ROp
	case Sqr:
		p.ROp *= p.ROp
	default:
		return errors.New("invalid function")
	}
	return nil
}

// GetLOpAndRes возвращает левый операнд.
func (p *TProc[T]) GetLOpAndRes() T {
	return p.LOpAndRes
}

// GetROp возвращает правый операнд.
func (p *TProc[T]) GetROp() T {
	return p.ROp
}

// GetOperation возвращает текущее состояние операции.
func (p *TProc[T]) GetOperation() TOperation {
	return p.Operation
}
