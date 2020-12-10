package main

import (
	"fmt"

	"github.com/gouarfig/adventofcode/helpers"
)

func main() {
	lines := helpers.GetLines("input.txt")
	var totalCode, totalValues, totalEncoded int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		code, values, encoded := calculateString(string(line))
		totalCode += code
		totalValues += values
		totalEncoded += encoded
	}
	fmt.Printf("1. total code=%d, values=%d, result=%d\n", totalCode, totalValues, totalCode-totalValues)
	fmt.Printf("2. total encoded=%d, code=%d, result=%d\n", totalEncoded, totalCode, totalEncoded-totalCode)
}

func calculateString(line string) (int, int, int) {
	var code, values, encoded int
	escape := false
	for i := 0; i < len(line); i++ {
		if !escape {
			switch line[i] {
			case '\\':
				code++
				encoded += 2
				escape = true

			case '"':
				code++
				encoded += 2

			default:
				code++
				values++
				encoded++
			}
			continue
		}
		// right after escape character
		switch line[i] {
		case '"', '\\':
			code++
			values++
			encoded += 2

		case 'x':
			code += 3
			encoded += 3
			values++
			i += 2

		default:
			code++
			values++
			encoded++
		}
		escape = false
	}
	encoded += 2
	return code, values, encoded
}
