package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemory_AllCases(t *testing.T) {
	mem := NewMemory()

	// Память пуста
	require.Nil(t, mem.Read())

	val1, _ := NewFractionNumber(1, 2)
	val2, _ := NewFractionNumber(1, 4)

	// Save & Read
	mem.Save(val1)
	read := mem.Read()
	require.Equal(t, "1/2", read.String())

	// Проверка, что Read() возвращает копию, не влияет на оригинал
	read.(*FractionNumber).numerator = 99
	require.Equal(t, "1/2", mem.Read().String())

	// Add к непустой памяти
	mem.Add(val2) // 1/2 + 1/4 = 3/4
	require.Equal(t, "3/4", mem.Read().String())

	// Очистка
	mem.Clear()
	require.Nil(t, mem.Read())

	// Add к пустой памяти
	mem.Add(val2)
	require.Equal(t, "1/4", mem.Read().String())
}
