package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Конвертер p1 → p2")

	buildUI(w)

	w.ShowAndRun()
}
