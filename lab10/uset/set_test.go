package uset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTSet(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		set := NewSet[int]()
		assert.True(t, set.IsEmpty())

		set.Add(1)
		assert.False(t, set.IsEmpty())
		assert.True(t, set.Contains(1))

		set.Remove(1)
		assert.False(t, set.Contains(1))
		assert.True(t, set.IsEmpty())

		set.Add(2)
		set.Clear()
		assert.True(t, set.IsEmpty())
	})

	t.Run("Union operation", func(t *testing.T) {
		set1 := NewSet[int]()
		set1.Add(1)
		set1.Add(2)

		set2 := NewSet[int]()
		set2.Add(3)
		set2.Add(2)

		unionSet := set1.Union(set2)
		assert.ElementsMatch(t, []int{1, 2, 3}, unionSet.Elements())
	})

	t.Run("Difference operation", func(t *testing.T) {
		set1 := NewSet[int]()
		set1.Add(1)
		set1.Add(2)
		set1.Add(3)

		set2 := NewSet[int]()
		set2.Add(2)

		diffSet := set1.Difference(set2)
		assert.ElementsMatch(t, []int{1, 3}, diffSet.Elements())
	})

	t.Run("Intersection operation", func(t *testing.T) {
		set1 := NewSet[int]()
		set1.Add(1)
		set1.Add(2)
		set1.Add(3)

		set2 := NewSet[int]()
		set2.Add(2)
		set2.Add(3)
		set2.Add(4)

		intersectionSet := set1.Intersection(set2)
		assert.ElementsMatch(t, []int{2, 3}, intersectionSet.Elements())
	})

	t.Run("Size and Elements", func(t *testing.T) {
		set := NewSet[int]()
		set.Add(1)
		set.Add(2)
		set.Add(3)

		assert.Equal(t, 3, set.Size())
		assert.ElementsMatch(t, []int{1, 2, 3}, set.Elements())
	})
}
