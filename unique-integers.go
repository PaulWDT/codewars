package main

import (
	"fmt"
	"sort"
)

func unique(in []int) (out []int) {

	mymap := make(map[int]bool)

	for _, num := range in {
		if _, present := mymap[num]; !present { // only if "num" was not yet present in the map it will be appended to the list of unique values
			out = append(out, num)
			mymap[num] = true
		}
	}
	/*
		for num, _ := range mymap {
			out = append(out, num)
		}
	*/
	sort.Ints(out)
	return out
}

func main() {
	in := []int{1, 2, 99, 55, 4, 4, 3, 4, 6, 7, 8, 1, 55, 4, 4, 5, 2, 2}

	fmt.Println("unique elements in [", in, "] are [", unique(in), " !")

}
