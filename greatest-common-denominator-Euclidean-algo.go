// https://www.freecodecamp.org/news/euclidian-gcd-algorithm-greatest-common-divisor/
// https://www.youtube.com/watch?v=yHwneN6zJmU

// GGT = Größter Gemeinsamer Teiler OR GCD = Greatest Common Denominator : euclidian-gcd-algorithm NON-recursive
package main

import "fmt"

func main() {
	fmt.Println("GCD of 750,150= (150)", gcd(750, 150))
	fmt.Println("GCD of 252,105= (21)", gcd(252, 105))
	fmt.Println("GCD of 33,12= (3)", gcd(33, 12))
}

func gcd(a, b int) int {

	for r := a % b; r > 0; r = a % b {
		a, b = b, r
	}
	return b
}
