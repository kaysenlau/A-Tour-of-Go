package main

import (
	"fmt"
	"math"
)

func Sqrt(start float64) float64 {
	for guess := 1.0; ; {
		state := guess
		guess -= (guess*guess - start) / (2 * guess)

		if float32(state) == float32(guess) {
			return guess
		}
	}
}

func main() {
	fmt.Println(Sqrt(6.0) == math.Sqrt(6.0)) // true
	fmt.Println(Sqrt(2) == math.Sqrt(2))     // true
	fmt.Println(Sqrt(866) == math.Sqrt(866)) // true
}
