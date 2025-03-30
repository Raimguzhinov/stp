package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strings"
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
	LabelFracSep  = "/"
)

type CalcMode int

const (
	ModeFraction CalcMode = iota
	ModeTPNumber
	ModeComplex
)

var modeNames = []string{"Дроби", "p-числа", "Комплексные"}
var currentMode = ModeFraction

func InitUI(w fyne.Window) {
	ctrl := NewControlUnit()

	display := widget.NewLabelWithStyle("0", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true, Bold: true})
	display.Wrapping = fyne.TextTruncate
	displayScroll := container.NewHScroll(display)
	decimalLabel := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Italic: true})

	var historyWin fyne.Window
	historyBtn := widget.NewButtonWithIcon("", theme.HistoryIcon(), func() {
		openHistoryWindow(w, historyWin, ctrl)
	})
	copyBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(ctrl.CopyExpression())
	})
	pasteBtn := widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {
		text := w.Clipboard().Content()
		result := ctrl.PasteExpression(text)
		if result == "Ошибка" {
			dialog.ShowError(fmt.Errorf("некорректное выражение: %s", text), w)
			return
		}
		updateDisplay(display, decimalLabel, result)
	})

	var fractionUI, tpUI, complexUI *fyne.Container

	baseLabel := widget.NewLabel("Основание:")
	tpBase := widget.NewSlider(2, 16)
	tpBase.Step = 1
	tpBase.Value = 10
	tpBase.Orientation = widget.Horizontal

	precLabel := widget.NewLabel("Точность:")
	tpPrec := widget.NewSlider(0, 10)
	tpPrec.Step = 1
	tpPrec.Value = 0
	tpPrec.Orientation = widget.Horizontal

	tpControls := container.NewVBox(
		baseLabel, tpBase,
		precLabel, tpPrec,
	)
	tpControls.Hide() // по умолчанию скрыт

	// ====== Общие управляющие кнопки (MS/MR/M+/MC) ======
	msRow := container.NewGridWithColumns(4,
		widget.NewButton(LabelMS, func() { ctrl.MemorySave() }),
		widget.NewButton(LabelMR, func() { updateDisplay(display, decimalLabel, ctrl.MemoryRead()) }),
		widget.NewButton(LabelMPlus, func() { ctrl.MemoryAdd() }),
		widget.NewButton(LabelMC, func() { ctrl.MemoryClear() }),
	)

	// ====== Адаптивные управляющие кнопки (C/← [+ Sqr/1/x]) ======
	var controlRowBox *fyne.Container
	updateControlRow := func() {
		var row *fyne.Container
		if currentMode == ModeFraction {
			row = container.NewGridWithColumns(4,
				widget.NewButton(LabelClear, func() {
					updateDisplay(display, decimalLabel, ctrl.Input(LabelClear))
				}),
				widget.NewButton(LabelBack, func() {
					updateDisplay(display, decimalLabel, ctrl.Input(LabelBack))
				}),
				widget.NewButton(LabelInverse, func() {
					updateDisplay(display, decimalLabel, ctrl.ApplyFunction(LabelInverse))
				}),
				widget.NewButton(LabelSqr, func() {
					updateDisplay(display, decimalLabel, ctrl.ApplyFunction(LabelSqr))
				}),
			)
		} else {
			row = container.NewGridWithColumns(2,
				widget.NewButton(LabelClear, func() {
					updateDisplay(display, decimalLabel, ctrl.Input(LabelClear))
				}),
				widget.NewButton(LabelBack, func() {
					updateDisplay(display, decimalLabel, ctrl.Input(LabelBack))
				}),
			)
		}
		controlRowBox.Objects = []fyne.CanvasObject{row}
		controlRowBox.Refresh()
	}
	controlRowBox = container.NewVBox()
	updateControlRow()

	modeSelect := widget.NewSelect(modeNames, func(sel string) {
		switch sel {
		case "Дроби":
			currentMode = ModeFraction
			fractionUI.Show()
			tpUI.Hide()
			tpControls.Hide()
			complexUI.Hide()
		case "p-числа":
			currentMode = ModeTPNumber
			fractionUI.Hide()
			tpUI.Show()
			tpControls.Show()
			complexUI.Hide()
		case "Комплексные":
			currentMode = ModeComplex
			fractionUI.Hide()
			tpUI.Hide()
			tpControls.Hide()
			complexUI.Show()
		}
		updateControlRow()
	})
	modeSelect.Selected = modeNames[0]

	labelRow := container.NewHBox(
		widget.NewLabelWithStyle("Режим:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		modeSelect,
		layout.NewSpacer(),
		copyBtn,
		pasteBtn,
		historyBtn,
	)

	displayContainer := container.NewVBox(
		labelRow,
		displayScroll,
		decimalLabel,
	)

	// ====== Fraction UI ======
	numButtons := []string{
		"7", "8", "9", LabelDivide,
		"4", "5", "6", LabelMultiply,
		"1", "2", "3", LabelMinus,
		LabelNegate, "0", LabelDot, LabelPlus,
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
	fractionUI = container.NewVBox(numGrid, bottomRow)

	// ====== TPNumber UI ======

	digitButtons := make([]*widget.Button, 16)
	digitRow := container.NewGridWithColumns(4)
	for i := 0; i < 16; i++ {
		lbl := fmt.Sprintf("%X", i)
		btn := widget.NewButton(lbl, func(d string) func() {
			return func() {
				updateDisplay(display, decimalLabel, ctrl.Input(d))
			}
		}(lbl))
		digitButtons[i] = btn
		digitRow.Add(btn)
	}

	updateDigitButtonStates := func(base int) {
		for i, btn := range digitButtons {
			if i < base {
				btn.Enable()
			} else {
				btn.Disable()
			}
		}
	}
	updateDigitButtonStates(int(tpBase.Value))
	tpBase.OnChanged = func(v float64) {
		base := int(v)
		baseLabel.SetText(fmt.Sprintf("Основание: %d", base))
		updateDigitButtonStates(base)
	}
	tpPrec.OnChanged = func(v float64) {
		precLabel.SetText(fmt.Sprintf("Точность: %d", int(v)))
	}
	tpOps := container.NewGridWithColumns(4)
	for _, op := range []string{LabelPlus, LabelMinus, LabelMultiply, LabelDivide} {
		btn := widget.NewButton(op, func(opr string) func() {
			return func() {
				updateDisplay(display, decimalLabel, ctrl.Input(opr))
			}
		}(op))
		btn.Importance = widget.HighImportance
		tpOps.Add(btn)
	}
	tpEqual := widget.NewButton(LabelEqual, func() {
		updateDisplay(display, decimalLabel, ctrl.Evaluate())
	})
	tpEqual.Importance = widget.HighImportance

	tpUI = container.NewVBox(
		digitRow,
		tpOps,
		tpEqual,
	)
	tpUI.Hide()

	// ====== Complex UI ======
	cplxGrid := container.NewGridWithColumns(4)
	for _, label := range []string{
		"7", "8", "9", LabelDivide,
		"4", "5", "6", LabelMultiply,
		"1", "2", "3", LabelMinus,
		"i", "0", LabelDot, LabelPlus,
	} {
		btn := widget.NewButton(label, func(lbl string) func() {
			return func() {
				updateDisplay(display, decimalLabel, ctrl.Input(lbl))
			}
		}(label))
		if label == LabelPlus || label == LabelMinus || label == LabelMultiply || label == LabelDivide {
			btn.Importance = widget.HighImportance
		}
		cplxGrid.Add(btn)
	}
	cplxEqual := widget.NewButton(LabelEqual, func() {
		updateDisplay(display, decimalLabel, ctrl.Evaluate())
	})
	cplxEqual.Importance = widget.HighImportance
	complexUI = container.NewVBox(cplxGrid, cplxEqual)
	complexUI.Hide()

	// ====== MAIN LAYOUT ======
	mainContent := container.NewVBox(
		displayContainer,
		layout.NewSpacer(),
		tpControls,
		controlRowBox,
		msRow,
		fractionUI,
		tpUI,
		complexUI,
	)

	w.Resize(fyne.NewSize(400, 600))
	w.SetFixedSize(true)
	w.SetContent(mainContent)
}

func updateDisplay(display *widget.Label, decimalLabel *widget.Label, text string) {
	display.SetText(text)
	if n, err := ParseNumber(text); err == nil {
		if f, ok := n.(*FractionNumber); ok && f.Denominator() != 1 {
			d := fmt.Sprintf("≈ %.10g", n.Float64())
			decimalLabel.SetText(d)
		} else {
			decimalLabel.SetText("")
		}
	} else {
		decimalLabel.SetText("")
	}
}

func openHistoryWindow(parent fyne.Window, historyWin fyne.Window, ctrl *ControlUnit) {
	history := ctrl.HistoryList()
	if len(history) == 0 {
		dialog.ShowInformation("История", "История пуста", parent)
		return
	}

	var items []fyne.CanvasObject
	for i, entry := range history {
		resultOnly := ""
		if parts := strings.Split(entry, "="); len(parts) == 2 {
			resultOnly = strings.TrimSpace(parts[1])
		}
		index := i

		label := widget.NewLabel(fmt.Sprintf("%d) %s", i+1, entry))
		copyBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
			if resultOnly != "" {
				parent.Clipboard().SetContent(resultOnly)
			}
		})
		delBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			ctrl.DeleteHistory(index)
			if historyWin != nil {
				historyWin.Close()
			}
			openHistoryWindow(parent, historyWin, ctrl)
		})
		row := container.NewBorder(nil, nil, nil, container.NewHBox(copyBtn, delBtn), label)
		items = append(items, row)
	}

	scroll := container.NewVScroll(container.NewVBox(items...))
	historyWin = fyne.CurrentApp().NewWindow("История вычислений")
	historyWin.Resize(fyne.NewSize(400, 500))
	historyWin.SetFixedSize(true)
	historyWin.SetContent(container.NewBorder(
		widget.NewLabelWithStyle("Последние операции:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		nil, nil, nil, scroll,
	))
	historyWin.Show()
}
