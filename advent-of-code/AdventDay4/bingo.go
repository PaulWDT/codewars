package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	//binSlice := make([]string, 0, 1000)
	//numBits := len(bitCounts)

	for scanner.Scan() {
		lineContent := scanner.Text()
		lineNum++
		fmt.Println("line : ", lineNum, " len : ", len(lineContent), " contains : ", lineContent)
		//binSlice = append(binSlice, lineContent)
	}

	fmt.Println("Total number of lines =", lineNum)
}
