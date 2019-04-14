package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1.0; i <= x; i++ {
		z -= (z*z - x) / (2*z)
		if z != x {
			continue
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(16))
}
