package main

import "fmt"

func Sqrt(x float64, precision float64) float64 {
	z := x / 2
	counter := 0
	for ; (z*z-x)*(z*z-x) > precision*precision; z -= (z*z - x) / (2 * z) {
		counter = counter + 1
	}
	fmt.Printf("Took %v tries\n", counter)
	return z
}

func main() {
	fmt.Println(Sqrt(999, 0.01))
}
