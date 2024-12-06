package ueditor

import (
	"testing"
)

func TestNewComplexEditor(t *testing.T) {
	editor := NewComplexEditor()
	if editor.value != 0 {
		t.Errorf("expected value to be 0, got %v", editor.value)
	}
}

func TestComplexIsZero(t *testing.T) {
	editor := NewComplexEditor()
	if !editor.ComplexIsZero() {
		t.Errorf("expected ComplexIsZero to return true")
	}

	editor.value = 1 + 2i
	if editor.ComplexIsZero() {
		t.Errorf("expected ComplexIsZero to return false")
	}
}

func TestAddSign(t *testing.T) {
	editor := NewComplexEditor()

	err := editor.AddSign(RealPart)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != 0 {
		t.Errorf("expected real part to remain 0, got %v", real(editor.value))
	}

	editor.value = 5 + 3i
	err = editor.AddSign(RealPart)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != -5 {
		t.Errorf("expected real part to be -5, got %v", real(editor.value))
	}

	err = editor.AddSign(ImagPart)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if imag(editor.value) != -3 {
		t.Errorf("expected imaginary part to be -3, got %v", imag(editor.value))
	}

	err = editor.AddSign("invalid")
	if err == nil {
		t.Errorf("expected error for invalid part type")
	}
}

func TestAddDigit(t *testing.T) {
	editor := NewComplexEditor()

	err := editor.AddDigit(RealPart, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != 1 {
		t.Errorf("expected real part to be 1, got %v", real(editor.value))
	}

	err = editor.AddDigit(ImagPart, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if imag(editor.value) != 2 {
		t.Errorf("expected imaginary part to be 2, got %v", imag(editor.value))
	}

	err = editor.AddDigit("invalid", 1)
	if err == nil {
		t.Errorf("expected error for invalid part type")
	}

	err = editor.AddDigit(RealPart, 10)
	if err == nil {
		t.Errorf("expected error for digit out of range")
	}
}

func TestAddZero(t *testing.T) {
	editor := NewComplexEditor()

	err := editor.AddZero(RealPart)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != 0 {
		t.Errorf("expected real part to remain 0, got %v", real(editor.value))
	}

	err = editor.AddZero(ImagPart)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if imag(editor.value) != 0 {
		t.Errorf("expected imaginary part to remain 0, got %v", imag(editor.value))
	}

	err = editor.AddZero("invalid")
	if err == nil {
		t.Errorf("expected error for invalid part type")
	}
}

func TestBackspace(t *testing.T) {
	editor := NewComplexEditor()
	editor.value = 123 + 456i

	err := editor.Backspace(RealPart)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != 12 {
		t.Errorf("expected real part to be 12, got %v", real(editor.value))
	}

	err = editor.Backspace(ImagPart)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if imag(editor.value) != 45 {
		t.Errorf("expected imaginary part to be 45, got %v", imag(editor.value))
	}

	err = editor.Backspace("invalid")
	if err == nil {
		t.Errorf("expected error for invalid part type")
	}
}

func TestClear(t *testing.T) {
	editor := NewComplexEditor()
	editor.value = 123 + 456i
	editor.Clear()
	if editor.value != 0 {
		t.Errorf("expected value to be 0, got %v", editor.value)
	}
}

func TestWriteString(t *testing.T) {
	editor := NewComplexEditor()

	err := editor.WriteString("5.5 + i4.4")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != 5.5 || imag(editor.value) != 4.4 {
		t.Errorf("expected value to be 5.5 + 4.4i, got %v", editor.value)
	}

	err = editor.WriteString("invalid")
	if err == nil {
		t.Errorf("expected error for invalid input")
	}
}

func TestEdit(t *testing.T) {
	editor := NewComplexEditor()

	_, err := editor.Edit(AddZeroToReal)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != 0 {
		t.Errorf("expected real part to remain 0, got %v", real(editor.value))
	}

	_, err = editor.Edit(ToggleRealSign)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if real(editor.value) != 0 {
		t.Errorf("expected real part to remain 0, got %v", real(editor.value))
	}

	_, err = editor.Edit(ToggleImagSign)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if imag(editor.value) != 0 {
		t.Errorf("expected imaginary part to remain 0, got %v", imag(editor.value))
	}

	_, err = editor.Edit(ClearEditor)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if editor.value != 0 {
		t.Errorf("expected value to be 0, got %v", editor.value)
	}

	_, err = editor.Edit(999)
	if err == nil {
		t.Errorf("expected error for invalid command")
	}
}
