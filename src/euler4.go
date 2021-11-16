package src

import (
	"fmt"
	"strconv"
)

func Euler4(numDigits int) []int {
	res := []int{1, 1, 1}
	fmt.Printf("Euler4 %d %v \n", numDigits, res)

	largestPalindromeProduct := 1
	if numDigits == 2 {
		low := 10
		high := 100
		for i := low; i < high; i++ {
			for j := i; j < high; j++ {
				prod := i * j
				if IsPalindromeNumber(prod) {
					if prod > largestPalindromeProduct {
						largestPalindromeProduct = prod
						res[0] = i
						res[1] = j
						res[2] = prod
					}
				}
			}
		}
	}

	return res
}

func IsPalindrome(text string) bool {
	if len(text) == 0 {
		return true
	} else if len(text) == 1 {
		return true
	}
	first := text[0]
	last := text[len(text)-1]
	if first != last {
		return false
	}
	middle := text[1 : len(text)-1]
	// println("isPal ", text, "->", middle)
	// isPal  9009 -> 00
	return IsPalindrome(middle)
}
func IsPalindromeNumber(number int) bool {
	text := strconv.Itoa(number)
	return IsPalindrome(text)
}
