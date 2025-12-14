package main

import (
	"fmt"
	"math"
)

// funct: x*e^x - sin(x)
func f(x float64) float64 {
	return x*math.Exp(x) - math.Sin(x)
}

func secant(x0, x1, tol float64, maxIter int) {
	fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n", "Iter", "x0", "x1", "x_new", "Err")
	fmt.Println("--------------------------------------------------")

	for i := 1; i <= maxIter; i++ {
		fx0 := f(x0)
		fx1 := f(x1)

		// x_new = x1 - f(x1) * (x1 - x0) / (f(x1) - f(x0))
		xNew := x1 - fx1*(x1-x0)/(fx1-fx0)

		err := math.Abs(xNew - x1)

		fmt.Printf("%-5d %-10.4f %-10.4f %-10.4f %-10.4f\n", i, x0, x1, xNew, err)

		// Check if done
		if err < tol {
			fmt.Println("--------------------------------------------------")
			fmt.Printf("Root found: %.4f\n", xNew)
			fmt.Printf("Steps taken: %d\n", i)
			return
		}

		x0 = x1
		x1 = xNew
	}

	fmt.Println("max iterations reached. unluck broo.")
}

func main() {
	secant(-2.0, -3.0, 0.001, 50)
}
