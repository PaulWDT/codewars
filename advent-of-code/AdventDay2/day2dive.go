package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	position, depth := 0, 0

	for scanner.Scan() {
		lineContent := scanner.Text()
		lineNum++
		fmt.Println(lineContent)
		words := strings.Split(lineContent, " ")
		value, _ := strconv.Atoi(words[1])
		switch words[0] {
		case "forward":
			position += value
		case "down":
			depth += value
		case "up":
			depth -= value
		default:
			fmt.Println("bad token ! (first word of line is not forward nor down nor up)")
		}

		/*		if err != nil {
				fmt.Println("Error : strconv.Atoi(", lineContent, ") failed !")
			} */

	}

	fmt.Println("Total number of steps =", lineNum)
	fmt.Println("position, depth =", position, depth)
	fmt.Println("multiply position * depth =", position*depth)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
