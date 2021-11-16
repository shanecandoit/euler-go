package src

func Euler3(num int64) int64 {
	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number
	// 600851475143 ?
	// println("Euler3", num)

	largestPrimeFactor := int64(1)

	factors := Factors(num)
	for f, _ := range factors {
		// println("fac", f)
		if f > largestPrimeFactor && IsPrime(f) {
			largestPrimeFactor = f
		}
	}

	// println("Euler3", num, "=", largestPrimeFactor)
	return int64(largestPrimeFactor)
}

// IsPrime a number is prime is it has no factors
func IsPrime(num int64) bool {
	if num == 1 {
		return true
	}
	facs := Factors(num)
	if len(facs) > 0 {
		return false
	}

	return true
}

// Factors the factors of num
func Factors(num int64) map[int64]int64 {
	original := num
	// dont add 1
	factors := map[int64]int64{}
	// fail early for 0 and smaller
	if num < 2 {
		return factors
	}
	if num%2 == 0 && num != 2 {
		//factors = append(factors, 2)
		factors[2] = 1
		num = num / 2
		for num%2 == 0 {
			factors[2] += 1
			num = num / 2
		}
	}
	for i := int64(3); i < num; i += 2 {
		for num%i == 0 {
			// println("fac", num, i)
			//factors = append(factors, i)
			c, ok := factors[i]
			if ok {
				factors[i] = int64(c + 1)
			} else {
				factors[i] = int64(1)
			}
			num = num / i
		}
	}
	// rem
	if original != num && num > 1 {
		c, ok := factors[num]
		if ok {
			factors[num] = int64(c + 1)
		} else {
			factors[num] = int64(1)
		}
	}

	return factors
}
