package main

import (
	"fmt"
)

func main() {
	fmt.Print(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""
	rows := make([]string, numRows)

	r := 0          // current row
	next_move := -1 // this variable indicates that we need to move forward or backward

	for i := 0; i < len(s); i++ {
		if r == 0 || r == numRows-1 { // change direction
			next_move *= -1
		}

		// append a character to the current row
		rows[r] += string(s[i])

		// move to the next row
		r += next_move
	}

	// merge strings of rows
	for i := 0; i < len(rows); i++ {
		res += rows[i]
	}

	return res
}
