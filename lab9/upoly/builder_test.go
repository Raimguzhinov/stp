package upoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolyBuilder(t *testing.T) {
	builderTests := []struct {
		name     string
		actions  func(*PolyBuilder) *TPoly
		expected *TPoly
	}{
		{
			"Build zero polynomial",
			func(b *PolyBuilder) *TPoly {
				return b.Build()
			},
			NewPoly(),
		},
		{
			"Build linear polynomial",
			func(b *PolyBuilder) *TPoly {
				return b.AddMember(3, 1).Build()
			},
			func() *TPoly {
				p := NewPoly()
				p.AddMember(NewMember(3, 1))
				return p
			}(),
		},
		{
			"Build quadratic polynomial",
			func(b *PolyBuilder) *TPoly {
				return b.AddMember(2, 2).AddMember(3, 1).Build()
			},
			func() *TPoly {
				p := NewPoly()
				p.AddMember(NewMember(2, 2))
				p.AddMember(NewMember(3, 1))
				return p
			}(),
		},
	}

	for _, tt := range builderTests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewPolyBuilder()
			result := tt.actions(b)
			assert.Equal(t, tt.expected.String(), result.String())
		})
	}
}
