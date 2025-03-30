package main

import (
	"fyne.io/fyne/v2/app"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

func main() {
	a := app.New()
	w := a.NewWindow("Универсальный калькулятор")
	log.SetLevel(log.DebugLevel)

	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Ошибка: %v\n", r)
			debug.PrintStack()
		}
	}()

	InitUI(w)
	w.ShowAndRun()
}
