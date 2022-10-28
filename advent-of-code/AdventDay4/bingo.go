/*
--- Day 4: Giant Squid ---

You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0

	8  2 23  4 24

21  9 14 16  7

	6 10  3 18  5
	1 12 20 15 19

	3 15  0  2 22
	9 18 13 17  5

19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5

	2  0 12  3  7

After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4

	8  2 23  4 24         9 18 13 17  5        10 16 15  9 19

21  9 14 16  7        19  8  7 25 23        18  8 23 26 20

	6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
	1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4

	8  2 23  4 24         9 18 13 17  5        10 16 15  9 19

21  9 14 16  7        19  8  7 25 23        18  8 23 26 20

	6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
	1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4

	8  2 23  4 24         9 18 13 17  5        10 16 15  9 19

21  9 14 16  7        19  8  7 25 23        18  8 23 26 20

	6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
	1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?
*/
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

	// fmt.Println("result for boards : ", boards)

	// solve bingo ...
	for _, num := range inValues {
		fmt.Print("\ncalling bingo = ", num, " : ")
		bingo(&boards, num)
	}
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

func bingo(boards *[]boardType, number int) {
	for boardIdx, board := range *boards {
		//board:=&boards[boardIdx]
		elimNumInBoard(&board, number) ////// PROBLEM - MUST CALL by Reference to MODIFY the original Slice-data !!!!
		win := checkBoardBingo(&board)
		if win {
			fmt.Println("Found the winning board at boardIndex = ", boardIdx, " board = ", board)
			panic("Stop on Bingo !")
		}
		if boardIdx == 20 {
			fmt.Println("display board at boardIndex = ", boardIdx, " board = ", board)
		}
	}
	return
}

func elimNumInBoard(board *boardType, number int) {
	for lineIdx, line := range board {
		for rowIdx, cell := range line {
			if cell == number {
				board[lineIdx][rowIdx] = 0 // if a cell matches the called "bingo number" then set it to zero
				fmt.Print("Match")
			}
		}
	}
}

func checkBoardBingo(board *boardType) bool {
	rowSum := []int{0, 0, 0, 0, 0}
	for lineIdx, line := range board {
		lineSum := 0
		for rowIdx, cell := range line {
			lineSum += cell
			rowSum[rowIdx] += cell
			if rowIdx == 4 && lineSum == 0 {
				fmt.Println("found a horizontal bingo :", board)
				return true
			}
			if lineIdx == 4 && rowSum[rowIdx] == 0 {
				fmt.Println("found a vertical bingo :", board)
				return true
			}
		}
		/*		if lineIdx==4 {
				for _,i := range rowSum {
					if i==0 {
						fmt.Println("found a vertical bingo :", board)
						return true
					}
				}
			} */

	}
	return false
}
