package src

// Euler1_Mult3or5 sum of multiples of 3 or 5
func Euler1_Mult3or5(top int) int {
	// If we list all the natural numbers
	// below 10 that are multiples of 3 or 5,
	// we get 3, 5, 6 and 9.
	// The sum of these multiples is 23.
	// Find the sum of all the multiples of 3 or 5 below 1000.
	// println("Euler1")

	var sum = 0
	var i = 1
	for i = 1; i < top; i++ {
		if i%3 == 0 || i%5 == 0 {

			sum += i
			// println("sum up to", i, "=", sum)
		}
	}
	// println("Euler1_Mult3or5", top, "=", sum)
	return sum
}
