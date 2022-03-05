package main

/*
In this little assignment you are given a string of space separated numbers,
and have to return the highest and lowest number.
Examples

HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"

Notes

    All numbers are valid Int32, no need to validate them.
    There will always be at least one number in the input string.
    Output string must be two numbers separated by a single space, and highest number is first.
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	tokens := strings.Fields(in)         // split string in words (spaces are separators)
	maxval, _ := strconv.Atoi(tokens[0]) // use first number to initialize
	minval := maxval

	for _, token := range tokens[1:] { // should start with second number
		if tmp, _ := strconv.Atoi(token); tmp > maxval {
			maxval = tmp
		} else if tmp < minval {
			minval = tmp
		}
	}

	return fmt.Sprintf("%d %d", maxval, minval) // always two return values even if equal or only one number in string
}

func main() {
	fmt.Println("HighAndLow(\"1 2 3 4 5\") returns ", HighAndLow("1 2 3 4 5"), " and SHOULD return \"5 1\"")
	fmt.Println("HighAndLow(\"1 2 -3 4 5\") returns ", HighAndLow("1 2 -3 4 5"), " and SHOULD return \"5 -3\"")
	fmt.Println("HighAndLow(\"1 9 3 4 -5\") returns ", HighAndLow("1 9 3 4 -5"), " and SHOULD return \"9 -5\"")
}
