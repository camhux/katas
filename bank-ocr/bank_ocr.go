package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var targetFigures [10]Figure

func init() {
	digitBytes, err := ioutil.ReadFile("fixtures/digits")
	if err != nil {
		log.Fatal("Failed to read digits file")
	}

	lines := PrepLinesFromBuffer(string(digitBytes), 0)

	for i := range targetFigures {
		targetFigures[i] = ReadFigure(lines, uint(i))
	}
}

func main() {
	// input := os.Args[1]

}

// Figure is a typealias for a [3]string representing a figure read from the input.
type Figure [3]string

// 9 figures expected per line
const lineFigures = 9

// ReadFigure reads the figure of the input at the zero-indexed offset from the start.
func ReadFigure(lines []string, offset uint) Figure {
	start := offset * 3
	end := start + 3
	ss0 := lines[0][start:end]
	ss1 := lines[1][start:end]
	ss2 := lines[2][start:end]

	return Figure{
		ss0,
		ss1,
		ss2,
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
func LineToNumerals(rawLine string) string {
	bytes := make([]byte, lineFigures)
	lines := PrepLinesFromBuffer(rawLine, 0)

	for i := range bytes {
		figure := ReadFigure(lines, uint(i))
		numeral, ok := FigureToNumeral(&figure)

		b := numeral
		if !ok {
			b = byte('?')
		}

		bytes[i] = b
	}

	return string(bytes)
}

func PrepLinesFromBuffer(buffer string, offset uint) []string {
	lines := strings.Split(buffer, "\n")
	return lines[offset : offset+3]
}
