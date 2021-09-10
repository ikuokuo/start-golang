package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	p := math.Pow10(-10)
	fmt.Printf("z=%g, p=%g\n", z, p)
	n := 0
	for {
		c := (z*z - x) / (2 * z)
		z -= c
		fmt.Printf(" [%d] z=%g\n", n, z)
		if math.Abs(c) <= p {
			break
		}
		n++
	}
	return z
}

func main() {
	x := 2.0
	fmt.Printf("Sqrt(%g)=%g\n", x, Sqrt(x))
	fmt.Printf("math.Sqrt(%g)=%g\n", x, math.Sqrt(x))
}
