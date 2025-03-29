package main

import (
	"fmt"
	"fyne.io/fyne/v2/dialog"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func buildUI(w fyne.Window) {
	w.Resize(fyne.NewSize(600, 800))

	ctrl := NewControl()

	input := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	output := widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	digitGrid := container.NewAdaptiveGrid(4)

	p1Slider := widget.NewSlider(2, 16)
	p2Slider := widget.NewSlider(2, 16)

	p1Select := widget.NewSelect([]string{}, nil)
	p2Select := widget.NewSelect([]string{}, nil)
	for i := 2; i <= 16; i++ {
		s := fmt.Sprintf("%d", i)
		p1Select.Options = append(p1Select.Options, s)
		p2Select.Options = append(p2Select.Options, s)
	}
	p1Select.SetSelected("10")
	p2Select.SetSelected("16")
	p1Slider.SetValue(10)
	p2Slider.SetValue(16)

	p1Select.OnChanged = func(val string) {
		if p, err := strconv.Atoi(val); err == nil {
			p1Slider.SetValue(float64(p))
			ctrl.Pin = p
			input.SetText(ctrl.Clear())
			output.SetText("0")
			updateButtons(digitGrid, ctrl.Pin)
		}
	}
	p2Select.OnChanged = func(val string) {
		if p, err := strconv.Atoi(val); err == nil {
			p2Slider.SetValue(float64(p))
			ctrl.Pout = p
			output.SetText(ctrl.Convert())
		}
	}

	p1Slider.OnChanged = func(v float64) {
		p1Select.SetSelected(fmt.Sprintf("%d", int(v)))
		ctrl.Pin = int(v)
		input.SetText(ctrl.Clear())
		output.SetText("0")
		updateButtons(digitGrid, ctrl.Pin)
	}
	p2Slider.OnChanged = func(v float64) {
		p2Select.SetSelected(fmt.Sprintf("%d", int(v)))
		ctrl.Pout = int(v)
		output.SetText(ctrl.Convert())
	}

	baseSelector := func(label string, slider *widget.Slider, selectWidget *widget.Select) fyne.CanvasObject {
		return container.NewVBox(
			widget.NewLabelWithStyle(label, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			slider, selectWidget,
			//container.NewHBox(slider, selectWidget),
		)
	}

	// === Цифровая панель ===
	for i := 0; i <= 15; i++ {
		val := i
		text := strings.ToUpper(strconv.FormatInt(int64(val), 16))
		btn := widget.NewButton(text, func() {
			if val >= ctrl.Pin {
				return
			}
			if ctrl.St == StateConverted {
				input.SetText(ctrl.Clear())
			}
			input.SetText(ctrl.DoEdit(val))
			output.SetText("0")
		})
		digitGrid.Add(btn)
	}
	updateButtons(digitGrid, ctrl.Pin)

	btnDot := widget.NewButton(".", func() {
		input.SetText(ctrl.DoEdit(16))
	})
	btnMinus := widget.NewButton("−", func() {
		input.SetText(ctrl.DoEdit(19))
	})
	btnBack := widget.NewButton("←", func() {
		input.SetText(ctrl.DoEdit(17))
	})
	btnClear := widget.NewButton("C", func() {
		input.SetText(ctrl.DoEdit(18))
		output.SetText("0")
	})
	btnConvert := widget.NewButton("Преобразовать", func() {
		output.SetText(ctrl.Convert())
	})
	btnSwap := widget.NewButton("↔", func() {
		p1 := ctrl.Pin
		p2 := ctrl.Pout
		ctrl.Pin = p2
		ctrl.Pout = p1
		p1Select.SetSelected(strconv.Itoa(ctrl.Pin))
		p2Select.SetSelected(strconv.Itoa(ctrl.Pout))
		p1Slider.SetValue(float64(ctrl.Pin))
		p2Slider.SetValue(float64(ctrl.Pout))
		output.SetText(ctrl.Convert())
		updateButtons(digitGrid, ctrl.Pin)
	})
	btnHistory := widget.NewButton("История", func() {
		showHistoryWindow(w, ctrl)
	})
	buttonsGrid := container.NewGridWithColumns(3,
		btnDot, btnMinus, btnBack,
		btnClear, btnSwap, btnConvert,
		btnHistory,
	)
	digitsPanel := container.NewVBox(
		digitGrid,
		buttonsGrid,
	)

	// === Системные кнопки (компактно, справа) ===

	content := container.NewVBox(
		widget.NewLabelWithStyle("Число (p1):", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		input,
		widget.NewLabelWithStyle("Результат (p2):", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		output,
		baseSelector("Основание p1:", p1Slider, p1Select),
		baseSelector("Основание p2:", p2Slider, p2Select),
		widget.NewLabelWithStyle("Цифры для ввода:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		digitsPanel,
	)

	centered := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(container.NewVBox(content)),
		layout.NewSpacer(),
	)

	w.SetContent(container.NewScroll(centered))

	// Клавиши
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyBackspace:
			input.SetText(ctrl.DoEdit(17))
		case fyne.KeyEscape, fyne.KeyDelete:
			input.SetText(ctrl.DoEdit(18))
			output.SetText("0")
		case fyne.KeyReturn, fyne.KeyEnter:
			output.SetText(ctrl.Convert())
		}
	})
	w.Canvas().SetOnTypedRune(func(r rune) {
		ch := strings.ToUpper(string(r))
		switch ch {
		case ".", ",":
			input.SetText(ctrl.DoEdit(16))
		case "-":
			input.SetText(ctrl.DoEdit(19))
		default:
			val := charToInt(r)
			if val >= 0 && val < ctrl.Pin {
				input.SetText(ctrl.DoEdit(val))
				output.SetText("0")
			}
		}
	})
}

func updateButtons(grid *fyne.Container, base int) {
	for _, obj := range grid.Objects {
		if btn, ok := obj.(*widget.Button); ok {
			val, err := strconv.ParseInt(btn.Text, 16, 32)
			if err == nil && int(val) < base {
				btn.Enable()
			} else {
				btn.Disable()
			}
		}
	}
}

func showHistoryWindow(parent fyne.Window, ctrl *Control) {
	count := ctrl.his.Count()
	if count == 0 {
		dialog.ShowInformation("История", "История пуста.", parent)
		return
	}

	records := ctrl.his.records

	list := widget.NewList(
		func() int {
			return len(records)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("...")
		},
		func(i int, o fyne.CanvasObject) {
			r := records[i]
			o.(*widget.Label).SetText(fmt.Sprintf(
				"%d) p%d: %s → p%d: %s", i+1, r.P1, r.NumberIn, r.P2, r.NumberOut))
		},
	)

	w := fyne.CurrentApp().NewWindow("История")
	w.Resize(fyne.NewSize(500, 400))
	w.SetContent(container.NewBorder(
		widget.NewLabelWithStyle("Последние преобразования:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		nil, nil, nil,
		container.NewMax(list),
	))
	w.Show()
}
