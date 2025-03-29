package main

import (
	"os"
	"testing"
)

func TestAddAndSort(t *testing.T) {
	pb := NewPhoneBook("test_contacts.csv")
	pb.Clear()

	pb.Add(Contact{"Сергей", "+7-999-111-11-11"})
	pb.Add(Contact{"Андрей", "+7-999-222-22-22"})
	pb.Add(Contact{"Борис", "+7-999-333-33-33"})

	expected := []string{"Андрей", "Борис", "Сергей"}
	for i, c := range pb.Contacts {
		if c.Name != expected[i] {
			t.Errorf("Expected %s at index %d, got %s", expected[i], i, c.Name)
		}
	}
}

func TestDelete(t *testing.T) {
	pb := NewPhoneBook("test_contacts.csv")
	pb.Clear()
	pb.Add(Contact{"Иван", "+7-123"})
	pb.Add(Contact{"Ольга", "+7-456"})
	pb.Delete(0)

	if len(pb.Contacts) != 1 || pb.Contacts[0].Name != "Ольга" {
		t.Errorf("Delete failed, got %+v", pb.Contacts)
	}
}

func TestFindByName(t *testing.T) {
	pb := NewPhoneBook("test_contacts.csv")
	pb.Clear()
	pb.Add(Contact{"Мария", "+7-999"})
	pb.Add(Contact{"Марк", "+7-888"})
	results := pb.FindByName("Мар")

	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}

func TestFormatPhoneNumber(t *testing.T) {
	num, ok := formatPhoneNumber("8(999)123-45-67")
	if !ok || num != "+8-999-123-45-67" {
		t.Errorf("Format failed: got %s, ok=%v", num, ok)
	}

	num2, ok2 := formatPhoneNumber("12345")
	if ok2 || num2 != "12345" {
		t.Errorf("Short number incorrectly formatted: %s", num2)
	}
}

func TestValidation(t *testing.T) {
	if isValidName("Иван123") {
		t.Error("Name validation failed: digits should be invalid")
	}
	if !isValidName("Иван") {
		t.Error("Valid name flagged as invalid")
	}
	if isValidNumber("+7ABC") {
		t.Error("Number with letters should be invalid")
	}
	if !isValidNumber("+7-999-123-45-67") {
		t.Error("Valid number flagged as invalid")
	}
}

func TestSaveAndLoad(t *testing.T) {
	tempFile := "temp_contacts.csv"
	pb := NewPhoneBook(tempFile)
	pb.Clear()
	pb.Add(Contact{"Тест", "+7-000-000-00-00"})
	err := pb.Save()
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	pb2 := NewPhoneBook(tempFile)
	if len(pb2.Contacts) != 1 || pb2.Contacts[0].Name != "Тест" {
		t.Errorf("Load failed, got %+v", pb2.Contacts)
	}
	os.Remove(tempFile)
}

func TestIgnoreInvalidCSV(t *testing.T) {
	f, _ := os.Create("bad.csv")
	f.WriteString("bad,data,too,many,fields\n")
	f.Close()

	pb := NewPhoneBook("bad.csv")
	err := pb.Load()
	if err != nil {
		t.Errorf("Should not error on malformed row, got: %v", err)
	}
	os.Remove("bad.csv")
}
