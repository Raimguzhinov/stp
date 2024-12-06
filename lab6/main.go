package main

import (
	"fmt"
	"stp/lab6/ueditor"
)

func main() {
	editor := ueditor.NewComplexEditor()
	fmt.Println("Начальное значение:", editor)

	// Добавляем цифры
	_ = editor.AddDigit(ueditor.RealPart, 1)
	_ = editor.AddDigit(ueditor.RealPart, 2)
	_ = editor.AddDigit(ueditor.ImagPart, 3)
	_ = editor.AddDigit(ueditor.ImagPart, 4)
	fmt.Println("После добавления цифр:", editor)

	// Меняем знаки
	_ = editor.AddSign(ueditor.RealPart)
	_ = editor.AddSign(ueditor.ImagPart)
	fmt.Println("После изменения знаков:", editor)

	// Удаляем последний символ
	_ = editor.Backspace(ueditor.RealPart)
	_ = editor.Backspace(ueditor.ImagPart)
	fmt.Println("После удаления символов:", editor)

	// Очищаем
	editor.Clear()
	fmt.Println("После очистки:", editor)

	// Читаем и записываем строки
	err := editor.WriteString("5.5 + i4.4")
	if err != nil {
		fmt.Println("Ошибка записи строки:", err)
	} else {
		fmt.Println("После записи строки:", editor.String())
	}
}
