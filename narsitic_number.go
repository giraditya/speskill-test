package main

import "math"

func countDigit(n int) int {
	var result int
	if n == 0 {
		return 0
	}
	result = 1 + countDigit(n/10)
	return result
}

func CheckNarsisticNumber(n int) bool {
	digitNumber := countDigit(n)
	dup := n
	sum := float64(0)

	for dup > 0 {
		sum += math.Pow(float64(dup%10), float64(digitNumber))
		dup /= 10
	}

	return float64(n) == sum
}
