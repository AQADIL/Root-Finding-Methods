package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x*math.Exp(x) - math.Sin(x)
}

func df(x float64) float64 {
	return math.Exp(x)*(1+x) - math.Cos(x)
}

func newton(x0, tol float64, maxIter int) {
	fmt.Printf("%-5s %-10s %-10s %-10s\n", "Iter", "x", "f(x)", "Err")
	fmt.Println("---------------------------------------")

	x := x0

	for i := 1; i <= maxIter; i++ {
		fx := f(x)
		dfx := df(x)

		if dfx == 0 {
			fmt.Println("error: derivative is zero. can't continue.")
			return
		}

		// formula
		xNew := x - (fx / dfx)
		err := math.Abs(xNew - x)

		fmt.Printf("%-5d %-10.4f %-10.4f %-10.4f\n", i, x, fx, err)

		if err < tol {
			fmt.Println("---------------------------------------")
			fmt.Printf("Root found: %.4f\n", xNew)
			fmt.Printf("Steps taken: %d\n", i)
			return
		}

		x = xNew
	}

	fmt.Println("max iterations reached. unluck broo.")
}

func main() {
	newton(-2.5, 0.001, 50)
}
