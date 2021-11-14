package main

import (
	"euler/src"
	"fmt"
)

func main() {
	fmt.Println("Euler")

	println(1, src.Euler1_Mult3or5(1_000))
	// sum up to 10 = 23
	// sum up to 999 = 233168

	// src.Euler2(10)
	// src.Euler2(4_000)
	println(2, src.Euler2(4_000_000))
	// Euler2  4000000 = 5242872
}
