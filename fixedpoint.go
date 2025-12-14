package main

import (
	"fmt"
	"math"
)

// we rewrote f(x)=0 into x = sin(x)/e^x
func g(x float64) float64 {
	return math.Sin(x) / math.Exp(x)
}

func fixedPoint(x0, tol float64, maxIter int) {
	fmt.Printf("%-5s %-10s %-10s %-10s\n", "Iter", "x_old", "x_new", "Err")
	fmt.Println("---------------------------------------")

	xPrev := x0

	for i := 1; i <= maxIter; i++ {
		xNew := g(xPrev) // calculate new x

		// check if number exploded (too huge)
		if math.IsInf(xNew, 0) || math.IsNaN(xNew) {
			fmt.Println("error: numbers got too crazy (diverged).")
			return
		}

		err := math.Abs(xNew - xPrev) // how much did it change?

		fmt.Printf("%-5d %-10.4f %-10.4f %-10.4f\n", i, xPrev, xNew, err)

		// Stopping condition
		if err < tol {
			fmt.Println("---------------------------------------")
			fmt.Printf("Root found: %.4f\n", xNew)
			fmt.Printf("Steps taken: %d\n", i)
			return
		}

		xPrev = xNew // update for next step
	}

	fmt.Println("max iterations reached. unluck broo.")
}

func main() {
	// Starting guess (x0).
	// Careful: Fixed-Point is moody.
	// Try 0.5 to find the root at 0.
	fixedPoint(0.5, 0.001, 50)
}
