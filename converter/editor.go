package main

import (
	"strings"
)

type Editor struct {
	number string
}

const (
	delimiter = "."
	zero      = "0"
	minus     = "-"
)

func (e *Editor) Number() string {
	return e.number
}

func (e *Editor) AddDigit(n int) string {
	if n < 0 || n > 15 {
		return e.number
	}
	e.number += string(intToChar(n))
	return e.number
}

func (e *Editor) AddZero() string {
	e.number += zero
	return e.number
}

func (e *Editor) AddDelim() string {
	if !strings.Contains(e.number, delimiter) {
		e.number += delimiter
	}
	return e.number
}

func (e *Editor) AddMinus() string {
	if strings.HasPrefix(e.number, minus) {
		e.number = strings.TrimPrefix(e.number, minus)
	} else {
		e.number = minus + e.number
	}
	return e.number
}

func (e *Editor) Backspace() string {
	if len(e.number) > 0 {
		e.number = e.number[:len(e.number)-1]
	}
	return e.number
}

func (e *Editor) Clear() string {
	e.number = ""
	return e.number
}

func (e *Editor) DoEdit(j int) string {
	switch j {
	case 0:
		return e.AddZero()
	case 16:
		return e.AddDelim()
	case 17:
		return e.Backspace()
	case 18:
		return e.Clear()
	case 19:
		return e.AddMinus()
	default:
		return e.AddDigit(j)
	}
}

func (e *Editor) Accuracy() int {
	parts := strings.Split(e.number, delimiter)
	if len(parts) == 2 {
		return len(parts[1])
	}
	return 0
}
