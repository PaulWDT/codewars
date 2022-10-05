package main

func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	xsout := []int{}
	ysout := []int{}

	swapmap := make(map[int]bool) // maps can only contain unique keys
	for _, val := range ns {
		swapmap[val] = true
	}

	swap := false
	for i, _ := range xs {
		if swapmap[i] == true { // if this index is contained in the map (don't care for order)
			swap = !swap // toggle the swapping on and off
		}
		if swap { // append genes to output chromosomes either crossover(swap) or straight
			xsout = append(xsout, ys[i])
			ysout = append(ysout, xs[i])
		} else {
			xsout = append(xsout, xs[i])
			ysout = append(ysout, ys[i])
		}
	}
	return xsout, ysout
}
