package main

import (
	"euler/src"
	"fmt"
)

func main() {
	fmt.Println("Euler")

	// println(1, src.Euler1_Mult3or5(1_000))
	// sum up to 10 = 23
	// sum up to 999 = 233168

	// println(2, src.Euler2(4_000_000))
	// src.Euler2(10)
	// src.Euler2(4_000)
	// Euler2  4000000 = 5242872

	// println(3, src.Euler3(13_195))
	// Euler3 13195 = 29
	println(3, src.Euler3(600_851_475_143))
	// Euler3 600851475143 = 6857

	fmt.Printf("%d %v\n", 4, src.Euler4(2))
	// 4 [91 99 9009]
}
