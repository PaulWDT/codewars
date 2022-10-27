package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type boardType [5][5]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	var inValues []int // sequence of input values ... to be fed into the bingo fields until "Win!"

	var boards []boardType
	var board boardType
	//	board := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}} // empty board as temp value
	boardNumber := 0
	var inBoardStrings [5]string
	inBoardStringIdx := 0

	for scanner.Scan() {
		lineContent := scanner.Text()
		lineNum++

		if lineNum == 1 { // first line is a list of bingo numbers
			inValues = parseInValues(lineContent)
			fmt.Println("input-values parsed into slice of ints:", inValues)
			continue
		}

		if len(lineContent) == 0 { // ignore empty lines
			continue
		}

		if inBoardStringIdx < 5 {
			inBoardStrings[inBoardStringIdx] = lineContent
			inBoardStringIdx++
		}
		if inBoardStringIdx == 5 {
			board = stringsToBoard(inBoardStrings)
			boards = append(boards, board)
			boardNumber++
			fmt.Println("board #", boardNumber, " is : ", board)
			inBoardStringIdx = 0 // restart idx for next board
		}

		//fmt.Println("line : ", lineNum, " len : ", len(lineContent), " contains : ", lineContent)
	}

	fmt.Println("Total number of lines =", lineNum) // 601 lines in input.txt - first line are the bingo inputs
	//num, _ := strconv.ParseInt("123456", 10, 32)
	fmt.Println("Total number of boards =", boardNumber) // gives 100 boards with 5 lines each + one empty line as separator

	fmt.Println("result for boards : ", boards)
}

func parseInValues(inStr string) (out []int) {
	inWords := strings.Split(inStr, ",")
	for _, word := range inWords {
		val, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}
		out = append(out, val)
	}
	return
}

func stringsToBoard(inStrings [5]string) (brd boardType) {
	for i := 0; i < 5; i++ {
		inIntStrings := strings.Fields(inStrings[i]) // any number of consecutive spaces treated as single separator !
		for idx, valStr := range inIntStrings {
			tmpInt, err := strconv.Atoi(valStr)
			brd[i][idx] = tmpInt
			if err != nil {
				panic(err)
			}
		}

	}
	return
}
