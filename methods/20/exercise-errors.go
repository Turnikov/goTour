package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	counter := 0
	for ; (z*z-x)*(z*z-x) > 1e-9; z -= (z*z - x) / (2 * z) {
		counter = counter + 1
	}
	fmt.Printf("Took %v tries\n", counter)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
