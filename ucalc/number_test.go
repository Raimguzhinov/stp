package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseNumber_Integer(t *testing.T) {
	n, err := ParseNumber("5")
	require.NoError(t, err)
	require.Equal(t, "5", n.String())

	n, err = ParseNumber("-7")
	require.NoError(t, err)
	require.Equal(t, "-7", n.String())
}

func TestParseNumber_Fraction(t *testing.T) {
	n, err := ParseNumber("2/3")
	require.NoError(t, err)
	require.Equal(t, "2/3", n.String())

	n, err = ParseNumber("-3/4")
	require.NoError(t, err)
	require.Equal(t, "-3/4", n.String())
}

func TestParseNumber_Decimal(t *testing.T) {
	n, err := ParseNumber("0.25")
	require.NoError(t, err)
	require.Equal(t, "1/4", n.String())

	n, err = ParseNumber("-1.5")
	require.NoError(t, err)
	require.Equal(t, "-3/2", n.String())

	_, err = ParseNumber("1.")
	require.Error(t, err)

	_, err = ParseNumber(".5")
	require.Error(t, err)
}

func TestParseNumber_Errors(t *testing.T) {
	_, err := ParseNumber("1/0")
	require.Error(t, err)

	_, err = ParseNumber("abc")
	require.Error(t, err)

	_, err = ParseNumber("1/2/3")
	require.Error(t, err)

	_, err = ParseNumber("1.2.3")
	require.Error(t, err)

	_, err = ParseNumber("1/")
	require.Error(t, err)

	_, err = ParseNumber("1/x")
	require.Error(t, err)

	_, err = ParseNumber("/2")
	require.Error(t, err)

	_, err = ParseNumber("1..2")
	require.Error(t, err)

	_, err = ParseNumber("-1/x")
	require.Error(t, err)

	_, err = ParseNumber("-1.2.3")
	require.Error(t, err)

	_, err = ParseNumber("-/3")
	require.Error(t, err)
}
