package upnumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTPNumber(t *testing.T) {
	t.Run("Constructor from number", func(t *testing.T) {
		num, err := NewTPNumberFromNumber(10.5, 10, 2)
		assert.NoError(t, err)
		assert.Equal(t, 10.5, num.GetValue())
		assert.Equal(t, 10, num.GetBase())
		assert.Equal(t, 2, num.GetPrecision())

		// Некорректное основание
		_, err = NewTPNumberFromNumber(10.5, 1, 2)
		assert.Error(t, err)

		// Некорректная точность
		_, err = NewTPNumberFromNumber(10.5, 10, -1)
		assert.Error(t, err)
	})

	t.Run("Constructor from string", func(t *testing.T) {
		num, err := NewTPNumberFromString("10.5", "10", "2")
		assert.NoError(t, err)
		assert.Equal(t, 10.5, num.GetValue())
		assert.Equal(t, 10, num.GetBase())
		assert.Equal(t, 2, num.GetPrecision())

		// Некорректные входные данные
		_, err = NewTPNumberFromString("invalid", "10", "2")
		assert.Error(t, err)

		_, err = NewTPNumberFromString("10.5", "invalid", "2")
		assert.Error(t, err)

		_, err = NewTPNumberFromString("10.5", "10", "invalid")
		assert.Error(t, err)
	})

	t.Run("Basic operations", func(t *testing.T) {
		num1, _ := NewTPNumberFromNumber(10, 10, 2)
		num2, _ := NewTPNumberFromNumber(5, 10, 2)

		// Сложение
		sum, err := num1.Add(num2)
		assert.NoError(t, err)
		assert.Equal(t, 15.0, sum.GetValue())

		// Вычитание
		diff, err := num1.Sub(num2)
		assert.NoError(t, err)
		assert.Equal(t, 5.0, diff.GetValue())

		// Умножение
		mul, err := num1.Mul(num2)
		assert.NoError(t, err)
		assert.Equal(t, 50.0, mul.GetValue())

		// Деление
		div, err := num1.Div(num2)
		assert.NoError(t, err)
		assert.Equal(t, 2.0, div.GetValue())

		// Деление на ноль
		numZero, _ := NewTPNumberFromNumber(0, 10, 2)
		_, err = num1.Div(numZero)
		assert.Error(t, err)
	})

	t.Run("Advanced operations", func(t *testing.T) {
		num, _ := NewTPNumberFromNumber(10, 10, 2)

		// Инверсия
		inv, err := num.Inverse()
		assert.NoError(t, err)
		assert.Equal(t, 0.1, inv.GetValue())

		// Инверсия нуля
		numZero, _ := NewTPNumberFromNumber(0, 10, 2)
		_, err = numZero.Inverse()
		assert.Error(t, err)

		// Квадрат
		sqr, err := num.Square()
		assert.NoError(t, err)
		assert.Equal(t, 100.0, sqr.GetValue())
	})

	t.Run("Getters and Setters", func(t *testing.T) {
		num, _ := NewTPNumberFromNumber(10.5, 10, 2)

		// Проверка базовых геттеров
		assert.Equal(t, 10.5, num.GetValue())
		assert.Equal(t, 10, num.GetBase())
		assert.Equal(t, 2, num.GetPrecision())

		// Изменение основания
		err := num.SetBase(16)
		assert.NoError(t, err)
		assert.Equal(t, 16, num.GetBase())

		err = num.SetBase(1)
		assert.Error(t, err)

		// Изменение точности
		err = num.SetPrecision(5)
		assert.NoError(t, err)
		assert.Equal(t, 5, num.GetPrecision())

		err = num.SetPrecision(-1)
		assert.Error(t, err)
	})

	t.Run("String representation", func(t *testing.T) {
		num, _ := NewTPNumberFromNumber(10.5, 10, 2)
		assert.Equal(t, "10.50", num.ToString())
	})
}
