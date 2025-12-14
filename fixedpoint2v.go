package main

import (
	"fmt"
	"math"
)

// x = -Pi - arcsin(x * e^x)
func g2(x float64) float64 {
	return -math.Pi - math.Asin(x*math.Exp(x))
}

func fixedPoint2v(x0, tol float64, maxIter int) {
	fmt.Printf("%-5s %-10s %-10s %-10s\n", "Iter", "x_old", "x_new", "Err")
	fmt.Println("---------------------------------------")

	xPrev := x0

	for i := 1; i <= maxIter; i++ {
		xNew := g2(xPrev) // calculate new x

		if math.IsNaN(xNew) {
			fmt.Println("error: math broke (NaN). Try closer x0.")
			return
		}

		err := math.Abs(xNew - xPrev)

		fmt.Printf("%-5d %-10.4f %-10.4f %-10.4f\n", i, xPrev, xNew, err)

		// Stopping condition
		if err < tol {
			fmt.Println("---------------------------------------")
			fmt.Printf("Root found: %.4f\n", xNew)
			fmt.Printf("Steps taken: %d\n", i)
			return
		}

		xPrev = xNew
	}

	fmt.Println("max iterations reached. unluck broo.")
}

func main() {

	fixedPoint2v(-2.5, 0.001, 50)
}
