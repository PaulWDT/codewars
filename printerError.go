/* In a factory a printer prints labels for boxes. For one kind of boxes the printer has to use colors which, for the sake of simplicity, are named with letters from a to m.

The colors used by the printer are recorded in a control string. For example a "good" control string would be aaabbbbhaijjjm meaning that the printer used three times color a, four times color b, one time color h then one time color a...

Sometimes there are problems: lack of colors, technical malfunction and a "bad" control string is produced e.g. aaaxbbbbyyhwawiwjjjwwm with letters not from a to m.

You have to write a function printer_error which given a string will return the error rate of the printer as a string representing a rational whose numerator is the number of errors and the denominator the length of the control string. Don't reduce this fraction to a simpler expression.

The string has a length greater or equal to one and contains only letters from a to z.
Examples:

s="aaabbbbhaijjjm"
printer_error(s) => "0/14"

s="aaaxbbbbyyhwawiwjjjwwm"
printer_error(s) => "8/22"

*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in := "aaabbbbhaijjjm"
	fmt.Println("PrinterError(\"", in, "\") = ", PrinterError(in), " should be 0/14")
	in = "aaaxbbbbyyhwawiwjjjwwm"
	fmt.Println("PrinterError(\"", in, "\") = ", PrinterError(in), " should be 8/22")
}

func PrinterError(in string) string {
	inLen := len(in)

	if inLen == 0 { // not required
		os.Exit(1)
	}
	in = strings.ToLower(in) // not required

	allowed := "abcdefghijklm"
	legal := 0
	for _, letter := range allowed { // range will yield runes
		legal += strings.Count(in, string(letter))
	}

	return fmt.Sprint(inLen-legal, "/", inLen)
}
