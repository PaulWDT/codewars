/* In 1978 the British Medical Journal reported on an outbreak of influenza at a British boarding school. There were 1000 students. The outbreak began with one infected student.

We want to study the spread of the disease through the population of this school. The total population may be divided into three:
the infected (i), those who have recovered (r), and those who are still susceptible (s) to get the disease.

We will study the disease on a period of tm days. One model of propagation uses 3 differential equations:

(1) s'(t) = -b * s(t) * i(t)
(2) i'(t) =  b * s(t) * i(t) - a * i(t)
(3) r'(t) =  a * i(t)

where s(t), i(t), r(t) are the susceptible, infected, recovered at time t and s'(t), i'(t), r'(t) the corresponding derivatives.
b and a are constants: b is representing a number of contacts which can spread the disease and a is a fraction of the infected that will recover.

We can transform equations (1), (2), (3) in finite differences (https://en.wikipedia.org/wiki/Finite_difference_method#Example:_ordinary_differential_equation)
(http://www.codewars.com/kata/56347fcfd086de8f11000014)

(I)    S[k+1] = S[k] - dt * b * S[k] * I[k]
(II)   I[k+1] = I[k] + dt * (b * S[k] * I[k] - a * I[k])
(III)  R[k+1] = R[k] + dt * I[k] *a

The interval [0, tm] will be divided in n small intervals of length dt = tm/n. Initial conditions here could be : S0 = 999, I0 = 1, R0 = 0
Whatever S0 and I0, R0 (number of recovered at time 0) is always 0.

The function epidemic will return the maximum number of infected as an integer (truncate to integer the result of max(I)).
Example:

tm = 14 ;n = 336 ;s0 = 996 ;i0 = 2 ;b = 0.00206 ;a = 0.41
epidemic(tm, n, s0, i0, b, a) --> 483

Notes:

    Keeping track of the values of susceptible, infected and recovered you can plot the solutions of the 3 differential equations. See an example below on the plot.

*/

package main

import (
	"fmt"
)

func Epidemic(tm, n, s0, i0 int, b, a float64) int { // tm=TimeMax, n=number-of-Intervals, i0=infected@t0, s0=people-susceptible@t0 (=population-infected)
	// b=rate-of-contact, a=rate-of-recovery (0 to <1)
	S := make([]float64, n+1) // S[t] Susceptible to be infected at time t
	I := make([]float64, n+1) // I(t) Infected at time t
	R := make([]float64, n+1) // R(t) Recovered at time t
	S[0] = float64(s0)
	I[0] = float64(i0)
	R[0] = 0.0
	var dt float64 = float64(tm) / float64(n)

	/* diese Version ist OK, aber verbesserbar ! (s.u.)
		max := I[0]

		for i := 0; i < n; i++ {
			S[i+1] = S[i] - dt*b*S[i]*I[i]
			I[i+1] = I[i] + dt*(b*S[i]*I[i]-a*I[i])
			R[i+1] = R[i] + dt*I[i]*a

			//max = math.Max(I[i+1], max)
			if I[i+1] > max {
				max = I[i+1]
			} else {
				break // no need to continue once the maximum is reached (if no increase anymore it CAN only decline from this point)
			}
		}
		fmt.Println("S[]=", S)
		fmt.Println("I[]=", I)
		fmt.Println("R[]=", R)

		return int(max)
	} */
	var deltaI float64

	for i := 0; i < n; i++ {
		S[i+1] = S[i] - dt*b*S[i]*I[i]

		deltaI = dt * (b*S[i]*I[i] - a*I[i])
		I[i+1] = I[i] + deltaI

		R[i+1] = R[i] + dt*I[i]*a

		if deltaI <= 0 {
			return int(I[i]) // no need to continue once the maximum is reached (deltaI <=0 then I[t] will decline from this point)
		}
	}
	return -1 // I(t) did not reach a maximum in our range of t
}

func main() {
	/* tm = 14 ;n = 336 ;s0 = 996 ;i0 = 2 ;b = 0.00206 ;a = 0.41
	   epidemic(tm, n, s0, i0, b, a) --> 483 */

	fmt.Println("Epidemic(14, 336, 996, 2, 0.00206, 0.41) = ", Epidemic(14, 336, 996, 2, 0.00206, 0.41))
	fmt.Println("and should be = 483")

}
