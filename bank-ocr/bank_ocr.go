package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

var targetFigures [9]Figure

func init() {
	digitBytes, err := ioutil.ReadFile("fixtures/digits")
	if err != nil {
		log.Fatal("Failed to read digits file")
	}

	digits := string(digitBytes)

	for i := range targetFigures {
		targetFigures[i] = ReadFigure(digits, uint(i))
	}
}

func main() {
	// input := os.Args[1]

}

// Figure is a typealias for a [3]string representing a figure read from the input.
type Figure [3]string

// input for OCR is a string of len (27 + 1 (newline)) * 3
const lineRuneLength = 27 + 1

// 9 figures expected per line
const lineFigures = 9

// ReadFigure reads the figure of the input at the zero-indexed offset from the start.
func ReadFigure(s string, offset uint) Figure {
	start := offset * 3
	end := start + 3
	ss1 := s[start:end]
	ss2 := s[start+lineRuneLength : end+lineRuneLength]
	ss3 := s[start+(lineRuneLength*2) : end+(lineRuneLength*2)]

	return Figure{
		ss1,
		ss2,
		ss3,
	}
}

// FigureToNumeral attempts to match a figure with the numeral digit it
// corresponds to. Return value `ok` indicates success or failure.
func FigureToNumeral(f *Figure) (res byte, ok bool) {
	for i, target := range targetFigures {
		if *f == target {
			ok = true
			res = byte(strconv.Itoa(i)[0])
			return
		}
	}
	ok = false
	return
}

// LineToNumerals reads a line (string where length is (27 + 1) * 3 )
// and returns a string of length 9 representing the best effort at
// reading the line into numerals
func LineToNumerals(line *string) string {
	bytes := make([]byte, lineFigures)

	for i := range bytes {
		figure := ReadFigure(*line, uint(i))
		numeral, ok := FigureToNumeral(&figure)

		b := numeral
		if !ok {
			b = byte('?')
		}

		bytes[i] = b
	}

	return string(bytes)
}
