package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x*math.Exp(x) - math.Sin(x)
}

func muller(x0, x1, x2, tol float64, maxIter int) {
	fmt.Printf("%-5s %-10s %-10s %-10s\n", "Iter", "x2 (curr)", "f(x2)", "Err")
	fmt.Println("---------------------------------------")

	for i := 1; i <= maxIter; i++ {

		f0 := f(x0)
		f1 := f(x1)
		f2 := f(x2)

		h1 := x1 - x0
		h2 := x2 - x1

		d1 := (f1 - f0) / h1
		d2 := (f2 - f1) / h2

		d := (d2 - d1) / (h2 + h1)

		// quadratic stuff
		b := d2 + h2*d
		D := math.Sqrt(b*b - 4*f2*d) // discriminant

		var E float64
		if math.Abs(b-D) < math.Abs(b+D) {
			E = b + D
		} else {
			E = b - D
		}

		// calculate correction (h)
		h := -2 * f2 / E
		xNew := x2 + h

		err := math.Abs(h) // error is the step size

		fmt.Printf("%-5d %-10.4f %-10.4f %-10.4f\n", i, x2, f2, err)

		// Check stopping criteria
		if err < tol {
			fmt.Println("---------------------------------------")
			fmt.Printf("Root found: %.4f\n", xNew)
			fmt.Printf("Steps taken: %d\n", i)
			return
		}

		x0 = x1
		x1 = x2
		x2 = xNew
	}

	fmt.Println("max iterations reached. unluck broo.")
}

func main() {

	muller(-2.5, -3.0, -3.5, 0.001, 50)
}
