package main

import "fmt"

func fibonacci() func() int {
	num, sum := 1, 0
	return func() int {
		num, sum = sum, num+sum
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
