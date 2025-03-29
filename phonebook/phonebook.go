// main.go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Contact struct {
	Name   string
	Number string
}

type PhoneBook struct {
	Contacts []Contact
	filePath string
}

func NewPhoneBook(filePath string) *PhoneBook {
	pb := &PhoneBook{filePath: filePath}
	pb.Load()
	return pb
}

func (pb *PhoneBook) Add(c Contact) {
	pb.Contacts = append(pb.Contacts, c)
	pb.sort()
}

func (pb *PhoneBook) Edit(index int, c Contact) {
	if index >= 0 && index < len(pb.Contacts) {
		pb.Contacts[index] = c
		pb.sort()
	}
}

func (pb *PhoneBook) Delete(index int) {
	if index >= 0 && index < len(pb.Contacts) {
		pb.Contacts = append(pb.Contacts[:index], pb.Contacts[index+1:]...)
	}
}

func (pb *PhoneBook) Clear() {
	pb.Contacts = []Contact{}
}

func (pb *PhoneBook) FindByName(name string) []Contact {
	var results []Contact
	for _, c := range pb.Contacts {
		if strings.Contains(strings.ToLower(c.Name), strings.ToLower(name)) {
			results = append(results, c)
		}
	}
	return results
}

func (pb *PhoneBook) sort() {
	sort.Slice(pb.Contacts, func(i, j int) bool {
		return strings.ToLower(pb.Contacts[i].Name) < strings.ToLower(pb.Contacts[j].Name)
	})
}

func (pb *PhoneBook) Save() error {
	f, err := os.Create(pb.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for _, c := range pb.Contacts {
		w.Write([]string{c.Name, c.Number})
	}
	w.Flush()
	return w.Error()
}

func (pb *PhoneBook) Load() error {
	f, err := os.Open(pb.filePath)
	if err != nil {
		return nil
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	pb.Contacts = nil
	for _, rec := range records {
		if len(rec) == 2 {
			pb.Contacts = append(pb.Contacts, Contact{Name: rec[0], Number: rec[1]})
		}
	}
	pb.sort()
	return nil
}

func formatPhoneNumber(input string) (string, bool) {
	digitsOnly := regexp.MustCompile(`\D`).ReplaceAllString(input, "")
	if len(digitsOnly) != 11 {
		return input, false
	}
	formatted := fmt.Sprintf("+%s-%s-%s-%s-%s",
		digitsOnly[0:1],
		digitsOnly[1:4],
		digitsOnly[4:7],
		digitsOnly[7:9],
		digitsOnly[9:11],
	)
	return formatted, true
}

func isValidName(name string) bool {
	return !regexp.MustCompile(`\d`).MatchString(name)
}

func isValidNumber(number string) bool {
	return !regexp.MustCompile(`[A-Za-zА-Яа-я]`).MatchString(number)
}

func main() {
	a := app.New()
	w := a.NewWindow("Телефонная книга")
	w.Resize(fyne.NewSize(400, 500))

	filePath := "contacts.csv"
	pb := NewPhoneBook(filePath)

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Имя")

	numberEntry := widget.NewEntry()
	numberEntry.SetPlaceHolder("Номер телефона")

	var selectedIndex int = -1

	list := widget.NewList(
		func() int { return len(pb.Contacts) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			if i >= 0 && i < len(pb.Contacts) {
				o.(*widget.Label).SetText(pb.Contacts[i].Name + ": " + pb.Contacts[i].Number)
			}
		},
	)

	list.OnSelected = func(id int) {
		selectedIndex = id
	}

	refresh := func() {
		list.Refresh()
	}

	addBtn := widget.NewButton("Добавить", func() {
		name := strings.TrimSpace(nameEntry.Text)
		numInput := strings.TrimSpace(numberEntry.Text)

		if name == "" || numInput == "" {
			dialog.ShowInformation("Ошибка", "Имя и номер обязательны", w)
			return
		}
		if !isValidName(name) {
			dialog.ShowError(fmt.Errorf("Имя не должно содержать цифры"), w)
			return
		}
		if !isValidNumber(numInput) {
			dialog.ShowError(fmt.Errorf("Номер не должен содержать буквы"), w)
			return
		}

		formatted, ok := formatPhoneNumber(numInput)

		confirmAdd := func(finalNum string) {
			pb.Add(Contact{Name: name, Number: finalNum})
			pb.Save()
			nameEntry.SetText("")
			numberEntry.SetText("")
			refresh()
		}

		if !ok {
			dialog.ShowConfirm("Нестандартный номер", "Формат номера отличается от +x-xxx-xxx-xx-xx. Сохранить как есть?",
				func(b bool) {
					if b {
						confirmAdd(numInput)
					}
				}, w)
		} else {
			confirmAdd(formatted)
		}
	})

	deleteBtn := widget.NewButton("Удалить", func() {
		if selectedIndex >= 0 && selectedIndex < len(pb.Contacts) {
			pb.Delete(selectedIndex)
			pb.Save()
			refresh()
			selectedIndex = -1
		}
	})

	clearBtn := widget.NewButton("Очистить", func() {
		dialog.ShowConfirm("Подтверждение", "Очистить все записи?", func(b bool) {
			if b {
				pb.Clear()
				pb.Save()
				refresh()
			}
		}, w)
	})

	chooseFileBtn := widget.NewButton("Выбрать CSV файл", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			if filepath.Ext(reader.URI().Name()) != ".csv" {
				dialog.ShowError(fmt.Errorf("Файл должен иметь расширение .csv"), w)
				return
			}
			pb.filePath = reader.URI().Path()
			err = pb.Load()
			if err != nil {
				dialog.ShowError(fmt.Errorf("Ошибка чтения файла: %v", err), w)
			} else {
				refresh()
			}
		}, w)
	})

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Поиск по имени")
	searchEntry.OnChanged = func(s string) {
		if s == "" {
			pb.Load()
		} else {
			pb.Contacts = pb.FindByName(s)
		}
		refresh()
	}

	toolbar := container.NewVBox(
		nameEntry,
		numberEntry,
		addBtn,
		deleteBtn,
		clearBtn,
		chooseFileBtn,
		searchEntry,
	)

	w.SetContent(container.NewBorder(toolbar, nil, nil, nil, list))
	w.ShowAndRun()
}
