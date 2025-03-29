package main

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type FractionEditor struct {
	buffer      string
	operand1    *FractionNumber
	operand2    *FractionNumber
	operation   string
	lastResult  *FractionNumber
	repeatOp    string
	repeatValue *FractionNumber
}

func NewFractionEditor() *FractionEditor {
	return &FractionEditor{}
}

func (f *FractionEditor) Input(char string) string {
	switch char {
	case LabelNegate:
		if strings.HasPrefix(f.buffer, "-") {
			f.buffer = strings.TrimPrefix(f.buffer, "-")
		} else {
			f.buffer = "-" + f.buffer
		}
	case LabelBack:
		if len(f.buffer) > 0 {
			f.buffer = f.buffer[:len(f.buffer)-1]
		}
	case LabelClear:
		f.clear()
	case LabelPlus, LabelMinus, LabelMultiply, LabelDivide:
		if f.buffer == "" {
			if f.lastResult != nil {
				f.operand1 = f.lastResult
				f.operation = char
			}
			return ""
		}
		frac, err := NewFractionFromString(f.buffer)
		if err != nil {
			return "Ошибка"
		}
		f.operand1 = frac
		f.operation = char
		f.buffer = ""
	case LabelFracSep:
		f.setOrReplaceChar(char, "1")
	case LabelDot:
		f.setOrReplaceChar(char, "0")
	default:
		f.buffer += char
	}
	return f.buffer
}

func (f *FractionEditor) setOrReplaceChar(char, autoSetFirstOperand string) {
	if !strings.Contains(f.buffer, char) {
		if f.buffer == "" {
			f.buffer += autoSetFirstOperand
		}
		f.buffer += char
	}
}

func (f *FractionEditor) clear() {
	log.Debug("очистка...")
	f.buffer = ""
	f.operand1 = nil
	f.operand2 = nil
	f.lastResult = nil
	f.operation = ""
	f.repeatOp = ""
	f.repeatValue = nil
}
