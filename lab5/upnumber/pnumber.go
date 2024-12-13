package upnumber

import (
	"errors"
	"strconv"
	"strings"
)

// TPNumber представляет р-ичное число.
type TPNumber struct {
	value     float64 // Действительное число
	base      int     // Основание системы счисления (2..16)
	precision int     // Точность представления числа (>= 0)
}

// NewTPNumberFromNumber создает р-ичное число из числа, основания и точности.
func NewTPNumberFromNumber(value float64, base, precision int) (*TPNumber, error) {
	if base < 2 || base > 16 {
		return nil, errors.New("base must be in range [2..16]")
	}
	if precision < 0 {
		return nil, errors.New("precision must be >= 0")
	}
	return &TPNumber{value: value, base: base, precision: precision}, nil
}

// NewTPNumberFromString создает р-ичное число из строкового представления.
func NewTPNumberFromString(valueStr, baseStr, precisionStr string) (*TPNumber, error) {
	base, err := strconv.Atoi(baseStr)
	if err != nil || base < 2 || base > 16 {
		return nil, errors.New("invalid base")
	}
	precision, err := strconv.Atoi(precisionStr)
	if err != nil || precision < 0 {
		return nil, errors.New("invalid precision")
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return nil, errors.New("invalid value")
	}
	return &TPNumber{value: value, base: base, precision: precision}, nil
}

// Copy создает копию р-ичного числа.
func (p *TPNumber) Copy() *TPNumber {
	return &TPNumber{value: p.value, base: p.base, precision: p.precision}
}

// Add складывает два р-ичных числа.
func (p *TPNumber) Add(other *TPNumber) (*TPNumber, error) {
	if p.base != other.base || p.precision != other.precision {
		return nil, errors.New("bases and precisions must match")
	}
	return NewTPNumberFromNumber(p.value+other.value, p.base, p.precision)
}

// Sub вычитает из числа другое р-ичное число.
func (p *TPNumber) Sub(other *TPNumber) (*TPNumber, error) {
	if p.base != other.base || p.precision != other.precision {
		return nil, errors.New("bases and precisions must match")
	}
	return NewTPNumberFromNumber(p.value-other.value, p.base, p.precision)
}

// Mul умножает два р-ичных числа.
func (p *TPNumber) Mul(other *TPNumber) (*TPNumber, error) {
	if p.base != other.base || p.precision != other.precision {
		return nil, errors.New("bases and precisions must match")
	}
	return NewTPNumberFromNumber(p.value*other.value, p.base, p.precision)
}

// Div делит число на другое р-ичное число.
func (p *TPNumber) Div(other *TPNumber) (*TPNumber, error) {
	if other.value == 0 {
		return nil, errors.New("division by zero")
	}
	if p.base != other.base || p.precision != other.precision {
		return nil, errors.New("bases and precisions must match")
	}
	return NewTPNumberFromNumber(p.value/other.value, p.base, p.precision)
}

// Inverse возвращает обратное число.
func (p *TPNumber) Inverse() (*TPNumber, error) {
	if p.value == 0 {
		return nil, errors.New("cannot invert zero")
	}
	return NewTPNumberFromNumber(1/p.value, p.base, p.precision)
}

// Square возвращает квадрат числа.
func (p *TPNumber) Square() (*TPNumber, error) {
	return NewTPNumberFromNumber(p.value*p.value, p.base, p.precision)
}

// GetValue возвращает значение числа.
func (p *TPNumber) GetValue() float64 {
	return p.value
}

// ToString возвращает строковое представление числа в системе счисления base.
func (p *TPNumber) ToString() string {
	result, err := formatFloatToBase(p.value, p.base, p.precision)
	if err != nil {
		return "Error: " + err.Error()
	}
	return result
}

// GetBase возвращает основание системы счисления.
func (p *TPNumber) GetBase() int {
	return p.base
}

// GetPrecision возвращает точность числа.
func (p *TPNumber) GetPrecision() int {
	return p.precision
}

// SetBase изменяет основание системы счисления.
func (p *TPNumber) SetBase(newBase int) error {
	if newBase < 2 || newBase > 16 {
		return errors.New("base must be in range [2..16]")
	}
	p.base = newBase
	return nil
}

// SetPrecision изменяет точность представления числа.
func (p *TPNumber) SetPrecision(newPrecision int) error {
	if newPrecision < 0 {
		return errors.New("precision must be >= 0")
	}
	p.precision = newPrecision
	return nil
}

// formatFloatToBase преобразует float64 в строку в заданной системе счисления с точностью.
func formatFloatToBase(value float64, base, precision int) (string, error) {
	if base < 2 || base > 16 {
		return "", errors.New("base must be in range [2..16]")
	}

	// Отделяем целую и дробную части
	intPart := int64(value)
	fracPart := value - float64(intPart)

	// Форматируем целую часть
	intStr := strconv.FormatInt(intPart, base)

	// Если точность 0, возвращаем только целую часть
	if precision == 0 {
		return intStr, nil
	}

	// Форматируем дробную часть
	var fracStr strings.Builder
	fracStr.WriteString(".")

	for i := 0; i < precision; i++ {
		fracPart *= float64(base)
		digit := int64(fracPart)
		fracStr.WriteString(strconv.FormatInt(digit, base))
		fracPart -= float64(digit)
	}

	// Объединяем целую и дробную часть
	return intStr + fracStr.String(), nil
}
