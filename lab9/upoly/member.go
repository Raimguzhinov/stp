package upoly

import (
	"fmt"
	"math"
)

// TMember представляет одночлен с целыми коэффициентом и степенью.
type TMember struct {
	Coeff  int // Коэффициент
	Degree int // Степень
}

// NewMember создает новый одночлен.
func NewMember(coeff, degree int) *TMember {
	return &TMember{
		Coeff:  coeff,
		Degree: degree,
	}
}

// Eval вычисляет значение одночлена в точке x.
func (m *TMember) Eval(x float64) float64 {
	return float64(m.Coeff) * math.Pow(x, float64(m.Degree))
}

// Diff возвращает производную одночлена.
func (m *TMember) Diff() *TMember {
	if m.Degree == 0 {
		return NewMember(0, 0)
	}
	return NewMember(m.Coeff*m.Degree, m.Degree-1)
}

// String возвращает строковое представление одночлена.
func (m *TMember) String() string {
	switch {
	case m.Coeff == 0:
		return ""
	case m.Degree == 0:
		return fmt.Sprintf("%d", m.Coeff)
	case m.Degree == 1:
		return fmt.Sprintf("%dx", m.Coeff)
	default:
		return fmt.Sprintf("%dx^%d", m.Coeff, m.Degree)
	}
}
