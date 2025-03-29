package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type ControlUnit struct {
	editor    *FractionEditor
	processor *Processor
	memory    *Memory
	history   *History
}

func NewControlUnit() *ControlUnit {
	return &ControlUnit{
		editor:    NewFractionEditor(),
		processor: NewProcessor(),
		memory:    NewMemory(),
		history:   NewHistory(),
	}
}

func (c *ControlUnit) Input(char string) string {
	return c.editor.Input(char)
}

func (c *ControlUnit) ApplyFunction(fn string) string {
	val, err := NewFractionFromString(c.editor.buffer)
	if err != nil {
		log.Error(err)
		return "Ошибка"
	}
	res, err := c.processor.ApplyUnary(fn, val)
	if err != nil {
		log.Error(err)
		return "Ошибка"
	}
	c.editor.buffer = res.String()
	return c.editor.buffer
}

func (c *ControlUnit) Evaluate() string {
	// Подготовка
	if c.editor.operand1 == nil && c.editor.lastResult != nil && c.editor.repeatOp != "" && c.editor.repeatValue != nil {
		c.editor.operand1 = c.editor.lastResult
		c.editor.operation = c.editor.repeatOp
		c.editor.operand2 = c.editor.repeatValue
	} else {
		if c.editor.operation == "" {
			return c.editor.buffer
		}
		val, err := NewFractionFromString(c.editor.buffer)
		if err != nil {
			log.Error(err)
			return "Ошибка"
		}
		c.editor.repeatValue = val
		c.editor.operand2 = val
	}

	// Вычисление
	res, err := c.processor.Execute(c.editor.operation, c.editor.operand1, c.editor.operand2)
	if err != nil {
		log.Error(err)
		return "Ошибка"
	}

	// Обновление
	c.history.Add(c.editor.operand1, c.editor.operation, c.editor.operand2, res.String())
	c.editor.lastResult = res
	c.editor.buffer = res.String()
	c.editor.operand1 = res
	c.editor.repeatOp = c.editor.operation
	c.editor.operation = ""

	return c.editor.buffer
}

func (c *ControlUnit) MemorySave() {
	if c.editor.lastResult != nil {
		c.memory.Save(c.editor.lastResult)
	}
}

func (c *ControlUnit) MemoryRead() string {
	if val := c.memory.Read(); val != nil {
		c.editor.buffer = val.String()
		return c.editor.buffer
	}
	return ""
}

func (c *ControlUnit) MemoryClear() {
	c.memory.Clear()
}

func (c *ControlUnit) MemoryAdd() {
	if c.editor.lastResult != nil {
		c.memory.Add(c.editor.lastResult)
	}
}

func (c *ControlUnit) HistoryList() []string {
	return c.history.Strings()
}

func (c *ControlUnit) DeleteHistory(index int) {
	c.history.Delete(index)
}

func (c *ControlUnit) CopyExpression() string {
	if c.editor.lastResult != nil {
		return c.editor.lastResult.String()
	}
	if c.editor.operand1 != nil && c.editor.operation != "" {
		return fmt.Sprintf("%s %s %s",
			c.editor.operand1.String(),
			c.editor.operation,
			c.editor.buffer,
		)
	}
	return c.editor.buffer
}

func (c *ControlUnit) PasteExpression(input string) string {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.TrimSuffix(input, "\n")

	// Если есть "=" — обрезаем результат
	if parts := strings.Split(input, "="); len(parts) == 2 {
		input = parts[0]
	}

	// Попытка распарсить как выражение (a op b)
	for _, op := range []string{LabelPlus, LabelMinus, LabelMultiply, LabelDivide} {
		if strings.Contains(input, op) {
			parts := strings.Split(input, op)
			if len(parts) != 2 {
				return "Ошибка"
			}
			a, err1 := NewFractionFromString(parts[0])
			b, err2 := NewFractionFromString(parts[1])
			if err1 != nil || err2 != nil {
				return "Ошибка"
			}

			res, err := c.processor.Execute(op, a, b)
			if err != nil {
				return "Ошибка"
			}

			c.editor.operand1 = a
			c.editor.operand2 = b
			c.editor.operation = op
			c.editor.lastResult = res
			c.editor.buffer = res.String()
			c.editor.repeatOp = op
			c.editor.repeatValue = b

			c.history.Add(a, op, b, res.String())
			return res.String()
		}
	}

	// Если нет оператора — пробуем вставить как отдельную дробь
	val, err := NewFractionFromString(input)
	if err != nil {
		return "Ошибка"
	}

	c.editor.buffer = val.String()
	return c.editor.buffer
}
