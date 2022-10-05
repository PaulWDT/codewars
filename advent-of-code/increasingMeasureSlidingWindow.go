package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(f.Stat())
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	var n, nMinus1, nMinus2, nMinus3 int
	numIncreased := 0
	for scanner.Scan() {
		lineContent := scanner.Text()
		lineNum++
		fmt.Println(lineContent)
		n, err = strconv.Atoi(lineContent)
		if err != nil {
			fmt.Println("Error : strconv.Atoi(", lineContent, ") failed !")
		}
		if lineNum > 3 && (n+nMinus1+nMinus2) > (nMinus1+nMinus2+nMinus3) {
			numIncreased++
		}
		nMinus3 = nMinus2
		nMinus2 = nMinus1
		nMinus1 = n
	}

	fmt.Println("Total number of measurements =", lineNum)
	fmt.Println("Number of increasing measurements =", numIncreased)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
