package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHistory_AddAndStrings(t *testing.T) {
	h := NewHistory()
	require.Len(t, h.Strings(), 0)

	a, _ := NewFractionNumber(1, 2)
	b, _ := NewFractionNumber(1, 4)
	h.Add(a, LabelPlus, b, "3/4")

	strs := h.Strings()
	require.Len(t, strs, 1)
	require.Equal(t, "1/2 + 1/4 = 3/4", strs[0])
}

func TestHistory_Delete(t *testing.T) {
	h := NewHistory()
	a, _ := NewFractionNumber(1, 2)
	b, _ := NewFractionNumber(1, 4)
	h.Add(a, LabelPlus, b, "3/4")
	h.Add(a, LabelMinus, b, "1/4")
	require.Len(t, h.Strings(), 2)

	h.Delete(0)
	require.Len(t, h.Strings(), 1)
	require.Contains(t, h.Strings()[0], "-")

	h.Delete(100) // должно быть безопасно
	require.Len(t, h.Strings(), 1)
}
