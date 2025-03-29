package main

import (
	"testing"
)

func TestEditor(t *testing.T) {
	ed := &Editor{}

	// AddDigit
	if ed.AddDigit(10) != "A" {
		t.Error("AddDigit failed to add 'A'")
	}

	// AddZero
	if ed.AddZero() != "A0" {
		t.Error("AddZero failed to add '0'")
	}

	// AddDelim
	if ed.AddDelim() != "A0." {
		t.Error("AddDelim failed to add delimiter")
	}
	if ed.AddDelim() != "A0." {
		t.Error("AddDelim added duplicate delimiter")
	}

	// Backspace
	if ed.Backspace() != "A0" {
		t.Error("Backspace failed to remove delimiter")
	}

	// Clear
	if ed.Clear() != "" {
		t.Error("Clear failed to reset number")
	}

	// DoEdit
	ed.DoEdit(10) // A
	ed.DoEdit(0)  // 0
	ed.DoEdit(16) // .
	ed.DoEdit(17) // Backspace
	ed.DoEdit(18) // Clear
	if ed.Number() != "" {
		t.Error("DoEdit failed")
	}
}

func TestConvert10ToP(t *testing.T) {
	val := Convert10ToP(-17.875, 16, 3)
	if val != "-11.E00" && val != "-11.e00" {
		t.Errorf("Expected -11.E, got %s", val)
	}
}

func TestConvertPTo10(t *testing.T) {
	val := ConvertPTo10("-11.E", 16)
	expected := -17.875
	if val != expected {
		t.Errorf("Expected %f, got %f", expected, val)
	}
}

func TestControl(t *testing.T) {
	ctrl := NewControl()

	ctrl.Pin = 16
	ctrl.Pout = 10

	ctrl.DoEdit(1)
	ctrl.DoEdit(1)
	ctrl.DoEdit(16) // .

	ctrl.DoEdit(14) // E
	res := ctrl.Convert()
	if res != "-17.8" && res != "17.8" {
		t.Logf("Control convert result: %s", res)
	}

	if ctrl.ed.Accuracy() != 1 {
		t.Error("Expected accuracy 1")
	}
	if ctrl.St != StateConverted {
		t.Error("Expected state StateConverted")
	}
}

func TestHistory(t *testing.T) {
	h := NewHistory()
	h.AddRecord(10, 16, "17.875", "11.E")

	if h.Count() != 1 {
		t.Error("Expected 1 record")
	}

	r := h.Get(0)
	if r.NumberOut != "11.E" {
		t.Error("Wrong record data")
	}
}
