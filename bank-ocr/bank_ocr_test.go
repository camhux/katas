package main

import (
	"io/ioutil"
	// "os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFigure(t *testing.T) {
	digitBytes, err := ioutil.ReadFile("testdata/sequence")
	if err != nil {
		t.Fatal("Failed to read sequence file")
	}

	input := string(digitBytes)

	expected1 := Figure{
		"   ",
		"  |",
		"  |",
	}

	actual1 := ReadFigure(input, 1)

	require.Equal(t, expected1, actual1)

	expected2 := Figure{
		" _ ",
		" _|",
		"|_ ",
	}

	actual2 := ReadFigure(input, 2)

	require.Equal(t, expected2, actual2)

	expected9 := Figure{
		" _ ",
		"|_|",
		" _|",
	}

	actual9 := ReadFigure(input, 8)

	require.Equal(t, expected9, actual9)
}

func TestFigureToNumeral(t *testing.T) {
	figureTwo := Figure{
		" _ ",
		" _|",
		"|_ ",
	}

	numeralTwo := byte('2')

	resultTwo, ok := FigureToNumeral(&figureTwo)

	require.True(t, ok, "Result was ok")

	require.Equal(t, numeralTwo, resultTwo,
		"Result was correct numeral as a string")

	figureFive := Figure{
		" _ ",
		"|_ ",
		" _|",
	}

	numeralFive := byte('5')

	resultFive, ok := FigureToNumeral(&figureFive)

	require.True(t, ok, "Result was ok")

	require.Equal(t, numeralFive, resultFive,
		"Result was correct numeral as a string")

	figureEight := Figure{
		" _ ",
		"|_|",
		"|_|",
	}

	numeralEight := byte('8')

	resultEight, ok := FigureToNumeral(&figureEight)

	require.True(t, ok, "Result was ok")

	require.Equal(t, numeralEight, resultEight,
		"Result was correct numeral as a string")
}

func TestLineToNumerals(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("testdata/account")
	if err != nil {
		t.Fatal("Failed to read testdata/account")
	}

	input := string(inputBytes)

	expected := "490067715"

	actual := LineToNumerals(&input)

	require.Equal(t, expected, actual)
}
