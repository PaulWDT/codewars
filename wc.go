package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	myargs := os.Args[1:]
	fmt.Println(myargs)
	if len(myargs) == 0 {
		fmt.Println(count(os.Stdin))
	} else {
		myfile, err := os.Open(myargs[0]) // first argument is filename to open, no other options yet

		if err != nil {
			fmt.Println(err)
			os.Exit(1) // returning something not zero means program failed
		}
		fmt.Println("Wordcount of file ", myargs[0], " = ", count(myfile))

	}

	os.Exit(0)
}

func count(r io.Reader) int {
	myscan := bufio.NewScanner(r)
	myscan.Split(bufio.ScanWords)
	var wc int

	for myscan.Scan() {
		wc++
	}
	return wc
}
