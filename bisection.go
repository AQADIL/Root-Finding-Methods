package main

import (
	"fmt"
	"math"
)

// 5th func: x*e^x - sin(x)
func f(x float64) float64 {
	return x*math.Exp(x) - math.Sin(x)
}

func bisection(a, b, tol float64, maxIter int) {
	// are the roots in this range?
	if f(a)*f(b) >= 0 {
		fmt.Println("error: pick other nums.")
		return
	}

	// table header
	fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n", "Iter", "a", "b", "Mid", "Err")
	fmt.Println("-------------------------------------------------------")

	for i := 1; i <= maxIter; i++ {
		c := (a + b) / 2.0     // midpoint
		err := math.Abs(b - a) // how big is the interval

		fmt.Printf("%-5d %-10.4f %-10.4f %-10.4f %-10.4f\n", i, a, b, c, err)

		// Stopping condition
		if err < tol || f(c) == 0 {
			fmt.Println("-------------------------------------------------------")
			fmt.Printf("Root found: %.4f\n", c)
			fmt.Printf("Steps taken: %d\n", i)
			return
		}

		// cut the interval in half
		if f(a)*f(c) < 0 {
			b = c // keep left side
		} else {
			a = c // keep right side
		}
	}

	fmt.Println("max iterations reached. unluck broo.")
}

func main() {

	bisection(-4, 1, 0.001, 50)
}
