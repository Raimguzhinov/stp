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
	LabelFracSep  = "/" // дробная черта
)

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

	labelRow := container.NewHBox(
		widget.NewLabelWithStyle("Ввод:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
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
	if f, err := NewFractionFromString(text); err == nil {
		if f.Denominator() != 1 {
			d := fmt.Sprintf("≈ %.10g", f.Float64())
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
			openHistoryWindow(parent, historyWin, ctrl) // перерисовка
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
