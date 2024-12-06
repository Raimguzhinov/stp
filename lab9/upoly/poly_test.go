package upoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTPoly(t *testing.T) {
	tests := []struct {
		name     string
		poly     *TPoly
		x        float64
		expected float64
	}{
		{"Zero polynomial", NewPoly(), 2, 0},
		{"Constant polynomial", func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(5, 0))
			return p
		}(), 2, 5},
		{"Linear polynomial", func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(3, 1))
			return p
		}(), 2, 6},
		{"Quadratic polynomial", func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(2, 2))
			p.AddMember(NewMember(3, 1))
			return p
		}(), 2, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.poly.Eval(tt.x))
		})
	}

	diffTests := []struct {
		name     string
		poly     *TPoly
		expected *TPoly
	}{
		{"Derivative of zero polynomial", NewPoly(), NewPoly()},
		{"Derivative of constant polynomial", func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(5, 0))
			return p
		}(), NewPoly()},
		{"Derivative of linear polynomial", func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(3, 1))
			return p
		}(), func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(3, 0))
			return p
		}()},
		{"Derivative of quadratic polynomial", func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(2, 2))
			p.AddMember(NewMember(3, 1))
			return p
		}(), func() *TPoly {
			p := NewPoly()
			p.AddMember(NewMember(4, 1))
			p.AddMember(NewMember(3, 0))
			return p
		}()},
	}

	for _, tt := range diffTests {
		t.Run(tt.name, func(t *testing.T) {
			expectedString := tt.expected.String()
			actualString := tt.poly.Diff().String()
			assert.Equal(t, expectedString, actualString)
		})
	}
}
