package upoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTMember(t *testing.T) {
	tests := []struct {
		name     string
		member   *TMember
		x        float64
		expected float64
	}{
		{"Zero member", NewMember(0, 0), 2, 0},
		{"Constant member", NewMember(5, 0), 2, 5},
		{"Linear member", NewMember(3, 1), 2, 6},
		{"Quadratic member", NewMember(2, 2), 2, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.member.Eval(tt.x))
		})
	}

	diffTests := []struct {
		name     string
		member   *TMember
		expected *TMember
	}{
		{"Derivative of constant", NewMember(5, 0), NewMember(0, 0)},
		{"Derivative of linear", NewMember(3, 1), NewMember(3, 0)},
		{"Derivative of quadratic", NewMember(2, 2), NewMember(4, 1)},
	}

	for _, tt := range diffTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.member.Diff())
		})
	}
}
