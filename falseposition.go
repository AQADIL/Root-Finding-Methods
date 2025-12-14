package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x*math.Exp(x) - math.Sin(x)
}

func falsePosition(a, b, tol float64, maxIter int) {

	if f(a)*f(b) >= 0 {
		fmt.Println("error: bad interval. pick other nums.")
		return
	}

	fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n", "Iter", "a", "b", "c", "Err")
	fmt.Println("--------------------------------------------------")

	var c float64     // The intersection point
	var cPrev float64 // To check error

	for i := 1; i <= maxIter; i++ {
		fa := f(a)
		fb := f(b)

		// it draws a line between (a, fa) and (b, fb)
		c = (a*fb - b*fa) / (fb - fa)

		fc := f(c)

		err := math.Abs(c - cPrev)

		fmt.Printf("%-5d %-10.4f %-10.4f %-10.4f %-10.4f\n", i, a, b, c, err)

		// Stopping criteria
		if math.Abs(fc) < tol || (i > 1 && err < tol) {
			fmt.Println("--------------------------------------------------")
			fmt.Printf("Root found: %.4f\n", c)
			fmt.Printf("Steps taken: %d\n", i)
			return
		}

		if fa*fc < 0 {
			b = c // Root is in left part
		} else {
			a = c // Root is in right part
		}

		cPrev = c
	}

	fmt.Println("max iterations reached. unluck broo.")
}

func main() {

	falsePosition(-4.0, -1.0, 0.001, 50)
}
