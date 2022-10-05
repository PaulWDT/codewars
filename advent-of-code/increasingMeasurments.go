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
	lastVal := 0
	numIncreased := 0
	for scanner.Scan() {
		lineContent := scanner.Text()
		lineNum++
		fmt.Println(lineContent)
		lineVal, err := strconv.Atoi(lineContent)
		if err != nil {
			fmt.Println("Error : strconv.Atoi(", lineContent, ") failed !")
		}
		if lineNum > 1 && lineVal > lastVal {
			numIncreased++
		}
		lastVal = lineVal
	}

	fmt.Println("Total number of measurements =", lineNum)
	fmt.Println("Number of increasing measurements =", numIncreased)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
