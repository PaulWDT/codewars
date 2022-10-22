package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const bits = 12

type node struct {
	left  *node
	right *node
	value int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	root := &node{}
	for scanner.Scan() {
		lines++
		l := scanner.Text()
		n, err := strconv.ParseInt(l, 2, 32)
		if err != nil {
			log.Fatalln(err.Error())
		}
		no := root
		for i := 0; i < bits; i++ {
			mask := int64(1 << (bits - 1 - i))
			if n&mask > 0 {
				if no.left == nil {
					no.left = &node{}
				}
				no.left.value += 1
				no = no.left
			} else {
				if no.right == nil {
					no.right = &node{}
				}
				no.right.value += 1
				no = no.right
			}
		}
		no.value = int(n)
	}
	no := root
	for {
		// left == 1, right==0
		if no.left == nil && no.right == nil {
			break
		}
		if no.left != nil && no.right == nil {
			no = no.left
			continue
		}
		if no.right != nil && no.left == nil {
			no = no.right
			continue
		}
		fmt.Printf("left: %v right: %v\n", no.left.value, no.right.value)
		if no.left.value >= no.right.value {
			no = no.left
		} else {
			no = no.right
		}
	}
	oxygen := no.value

	no = root
	for {
		// left == 1, right==0
		if no.left == nil && no.right == nil {
			break
		}
		if no.left != nil && no.right == nil {
			no = no.left
			continue
		}
		if no.right != nil && no.left == nil {
			no = no.right
			continue
		}

		fmt.Printf("left: %v right: %v\n", no.left.value, no.right.value)
		if no.left.value < no.right.value {
			no = no.left
		} else {
			no = no.right
		}
	}
	co2 := no.value

	fmt.Printf("answer: %v %v %v\n", oxygen, co2, oxygen*co2)
}
