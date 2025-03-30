package main

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type Editor struct {
	buffer      string
	operand1    Number
	operand2    Number
	operation   string
	lastResult  Number
	repeatOp    string
	repeatValue Number
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) Input(char string) string {
	switch char {
	case LabelNegate:
		if strings.HasPrefix(e.buffer, "-") {
			e.buffer = strings.TrimPrefix(e.buffer, "-")
		} else {
			e.buffer = "-" + e.buffer
		}
	case LabelBack:
		if len(e.buffer) > 0 {
			e.buffer = e.buffer[:len(e.buffer)-1]
		}
	case LabelClear:
		e.clear()
	case LabelPlus, LabelMinus, LabelMultiply, LabelDivide:
		if e.buffer == "" {
			if e.lastResult != nil {
				e.operand1 = e.lastResult
				e.operation = char
			}
			return ""
		}
		number, err := ParseNumber(e.buffer)
		if err != nil {
			log.Error(err)
			return "Ошибка"
		}
		e.operand1 = number
		e.operation = char
		e.buffer = ""
	case LabelFracSep:
		e.setOrReplaceChar(char, "1")
	case LabelDot:
		e.setOrReplaceChar(char, "0")
	default:
		e.buffer += char
	}
	return e.buffer
}

func (e *Editor) setOrReplaceChar(char, autoSetFirstOperand string) {
	if char == LabelDot && strings.Contains(e.buffer, LabelFracSep) {
		return
	}
	if char == LabelFracSep && strings.Contains(e.buffer, LabelDot) {
		return
	}
	if !strings.Contains(e.buffer, char) {
		if e.buffer == "" {
			e.buffer += autoSetFirstOperand
		}
		e.buffer += char
	}
}

func (e *Editor) clear() {
	log.Debug("очистка...")
	e.buffer = ""
	e.operand1 = nil
	e.operand2 = nil
	e.lastResult = nil
	e.operation = ""
	e.repeatOp = ""
	e.repeatValue = nil
}
