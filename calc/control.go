package main

import (
	log "github.com/sirupsen/logrus"
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
	c.editor.lastResult = res
	c.editor.buffer = res.String()
	c.editor.operand1 = res
	c.editor.repeatOp = c.editor.operation
	c.editor.operation = ""

	c.history.Add(c.editor.operand1, c.editor.repeatOp, c.editor.repeatValue, c.editor.buffer)
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
