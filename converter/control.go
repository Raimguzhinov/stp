package main

import (
	"math"
)

type ConverterState int

const (
	StateEditing ConverterState = iota
	StateConverted
)

type Control struct {
	Pin  int
	Pout int
	ed   *Editor
	St   ConverterState
	his  *History
}

func NewControl() *Control {
	return &Control{
		Pin:  10,
		Pout: 16,
		ed:   &Editor{},
		St:   StateEditing,
		his:  NewHistory(),
	}
}

func (c *Control) DoEdit(cmd int) string {
	c.St = StateEditing
	return c.ed.DoEdit(cmd)
}

func (c *Control) Convert() string {
	num10 := ConvertPTo10(c.ed.Number(), c.Pin)
	if math.Abs(num10) < 1e-12 {
		num10 = 0 // избавляемся от -0
	}
	acc := c.calcAccuracy()
	result := Convert10ToP(num10, c.Pout, acc)
	c.St = StateConverted
	c.his.AddRecord(c.Pin, c.Pout, c.ed.Number(), result)
	return result
}

func (c *Control) Clear() string {
	c.St = StateEditing
	return c.ed.Clear()
}

// Расчет точности результата
func (c *Control) calcAccuracy() int {
	acc := c.ed.Accuracy()
	return int(math.Round(float64(acc) * math.Log(float64(c.Pin)) / math.Log(float64(c.Pout))))
}
