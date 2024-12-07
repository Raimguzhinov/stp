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
	})

	t.Run("Constructor from string", func(t *testing.T) {
		num, err := NewTPNumberFromString("10.5", "10", "2")
		assert.NoError(t, err)
		assert.Equal(t, 10.5, num.GetValue())
		assert.Equal(t, 10, num.GetBase())
		assert.Equal(t, 2, num.GetPrecision())
	})

	t.Run("Operations", func(t *testing.T) {
		num1, _ := NewTPNumberFromNumber(10, 10, 2)
		num2, _ := NewTPNumberFromNumber(5, 10, 2)

		sum, err := num1.Add(num2)
		assert.NoError(t, err)
		assert.Equal(t, 15.0, sum.GetValue())

		diff, err := num1.Sub(num2)
		assert.NoError(t, err)
		assert.Equal(t, 5.0, diff.GetValue())

		mul, err := num1.Mul(num2)
		assert.NoError(t, err)
		assert.Equal(t, 50.0, mul.GetValue())

		div, err := num1.Div(num2)
		assert.NoError(t, err)
		assert.Equal(t, 2.0, div.GetValue())
	})
}
