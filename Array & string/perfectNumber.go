package main

import "math"

func checkPerfectNumber(n int) bool {
    if n <= 1 {
		return false
	}

	sum := 1

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}

	return sum == n
}