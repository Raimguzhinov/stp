package ueditor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type PartType string

const (
	RealPart PartType = "real"
	ImagPart PartType = "imag"
)

type CommandType int

const (
	AddZeroToReal  CommandType = iota // Добавить ноль в действительную часть
	ToggleRealSign                    // Переключить знак действительной части
	ToggleImagSign                    // Переключить знак мнимой части
	ClearEditor                       // Очистить редактор
)

type ComplexEditor struct {
	value complex128
}

func NewComplexEditor() *ComplexEditor {
	return &ComplexEditor{value: 0}
}

func (ce *ComplexEditor) String() string {
	realPart := fmt.Sprintf("%.2f", real(ce.value))
	imagPart := fmt.Sprintf("%.2f", imag(ce.value))
	if imag(ce.value) >= 0 {
		return fmt.Sprintf("%s + i%s", realPart, imagPart)
	}
	return fmt.Sprintf("%s - i%s", realPart, strings.TrimPrefix(imagPart, "-"))
}

func (ce *ComplexEditor) ComplexIsZero() bool {
	return ce.value == 0
}

func (ce *ComplexEditor) AddSign(part PartType) error {
	switch part {
	case RealPart:
		ce.value = complex(-real(ce.value), imag(ce.value))
	case ImagPart:
		ce.value = complex(real(ce.value), -imag(ce.value))
	default:
		return errors.New("invalid part type: must be 'real' or 'imag'")
	}
	return nil
}

func (ce *ComplexEditor) AddDigit(part PartType, digit int) error {
	if digit < 0 || digit > 9 {
		return errors.New("digit must be between 0 and 9")
	}
	switch part {
	case RealPart:
		ce.value = complex(real(ce.value)*10+float64(digit), imag(ce.value))
	case ImagPart:
		ce.value = complex(real(ce.value), imag(ce.value)*10+float64(digit))
	default:
		return errors.New("invalid part type: must be 'real' or 'imag'")
	}
	return nil
}

func (ce *ComplexEditor) AddZero(part PartType) error {
	return ce.AddDigit(part, 0)
}

func (ce *ComplexEditor) Backspace(part PartType) error {
	switch part {
	case RealPart:
		r := int(real(ce.value)) / 10
		ce.value = complex(float64(r), imag(ce.value))
	case ImagPart:
		i := int(imag(ce.value)) / 10
		ce.value = complex(real(ce.value), float64(i))
	default:
		return errors.New("invalid part type: must be 'real' or 'imag'")
	}
	return nil
}

func (ce *ComplexEditor) Clear() {
	ce.value = 0
}

func (ce *ComplexEditor) WriteString(input string) error {
	input = strings.ReplaceAll(input, " ", "")
	parts := strings.Split(input, "+i")
	if len(parts) != 2 {
		parts = strings.Split(input, "-i")
		if len(parts) != 2 {
			return errors.New("invalid input format")
		}
		parts[1] = "-" + parts[1]
	}

	realPart, err1 := strconv.ParseFloat(parts[0], 64)
	imagPart, err2 := strconv.ParseFloat(parts[1], 64)

	if err1 != nil || err2 != nil {
		return errors.New("failed to parse real or imaginary part")
	}

	ce.value = complex(realPart, imagPart)
	return nil
}

func (ce *ComplexEditor) Edit(command CommandType) (string, error) {
	switch command {
	case AddZeroToReal:
		_ = ce.AddZero(RealPart)
	case ToggleRealSign:
		_ = ce.AddSign(RealPart)
	case ToggleImagSign:
		_ = ce.AddSign(ImagPart)
	case ClearEditor:
		ce.Clear()
	default:
		return "", errors.New("invalid command")
	}
	return ce.String(), nil
}
