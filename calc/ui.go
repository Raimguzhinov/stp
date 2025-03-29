package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
)

const (
	LabelClear    = "C"
	LabelBack     = "←"
	LabelSqr      = "Sqr"
	LabelInverse  = "1/x"
	LabelMS       = "MS"
	LabelMR       = "MR"
	LabelMPlus    = "M+"
	LabelMC       = "MC"
	LabelNegate   = "±"
	LabelDot      = "."
	LabelDivide   = "÷"
	LabelMultiply = "×"
	LabelMinus    = "-"
	LabelPlus     = "+"
	LabelEqual    = "="
	LabelFracSep  = "/" // дробная черта
)

func InitUI(w fyne.Window) {
	ctrl := NewControlUnit()

	display := widget.NewLabelWithStyle("0", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true, Bold: true})
	display.Wrapping = fyne.TextTruncate
	displayScroll := container.NewHScroll(display)

	decimalLabel := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Italic: true})
	displayContainer := container.NewVBox(
		widget.NewLabelWithStyle("Ввод:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		displayScroll,
		decimalLabel,
	)

	numButtons := []string{
		"7", "8", "9", LabelDivide,
		"4", "5", "6", LabelMultiply,
		"1", "2", "3", LabelMinus,
		LabelNegate, "0", LabelDot, LabelPlus,
	}

	extras := []struct {
		label  string
		action func()
	}{
		{LabelClear, func() { updateDisplay(display, decimalLabel, ctrl.Input(LabelClear)) }},
		{LabelBack, func() { updateDisplay(display, decimalLabel, ctrl.Input(LabelBack)) }},
		{LabelSqr, func() { updateDisplay(display, decimalLabel, ctrl.ApplyFunction(LabelSqr)) }},
		{LabelInverse, func() { updateDisplay(display, decimalLabel, ctrl.ApplyFunction(LabelInverse)) }},
		{LabelMS, func() { ctrl.MemorySave() }},
		{LabelMR, func() { updateDisplay(display, decimalLabel, ctrl.MemoryRead()) }},
		{LabelMPlus, func() { ctrl.MemoryAdd() }},
		{LabelMC, func() { ctrl.MemoryClear() }},
	}

	extraRow := container.NewGridWithColumns(4)
	for _, e := range extras {
		btn := widget.NewButton(e.label, e.action)
		extraRow.Add(btn)
	}

	numGrid := container.NewGridWithColumns(4)
	for _, label := range numButtons {
		btn := widget.NewButton(label, func(lbl string) func() {
			return func() {
				updateDisplay(display, decimalLabel, ctrl.Input(lbl))
			}
		}(label))
		if label == LabelPlus || label == LabelMinus || label == LabelMultiply || label == LabelDivide {
			btn.Importance = widget.HighImportance
		}
		numGrid.Add(btn)
	}

	equalBtn := widget.NewButton(LabelEqual, func() {
		updateDisplay(display, decimalLabel, ctrl.Evaluate())
	})
	equalBtn.Importance = widget.HighImportance
	bottomRow := container.NewGridWithColumns(2,
		widget.NewButton(LabelFracSep, func() {
			updateDisplay(display, decimalLabel, ctrl.Input(LabelFracSep))
		}),
		equalBtn,
	)

	keyboard := container.NewVBox(extraRow, numGrid, bottomRow)

	mainContent := container.NewVBox(
		displayContainer,
		layout.NewSpacer(),
		keyboard,
	)

	w.Resize(fyne.NewSize(400, 600))
	w.SetFixedSize(true)
	w.SetContent(mainContent)
}

func updateDisplay(display *widget.Label, decimalLabel *widget.Label, text string) {
	display.SetText(text)
	f, err := NewFractionFromString(text)
	if err != nil {
		log.Error(err)
		decimalLabel.SetText("")
		return
	}
	if f.Denominator() != 1 {
		d := fmt.Sprintf("≈ %.10g", f.Float64())
		decimalLabel.SetText(d)
	} else {
		decimalLabel.SetText("")
	}
}
