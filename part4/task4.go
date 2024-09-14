// РАБОТАЕТ ТОЛЬКО ДЛЯ МАЛЕНЬКИХ ЗНАЧЕНИЙ

package main

import (
	"math"
)

func main() {

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func countDivisors(n int) int {
	count := 0

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i == i {
				count++
			} else {
				count += 2
			}
		}
	}

	return count
}

func simpleNumber(l, r int) int {
	result := []int{}

	for i := l; i <= r; i++ {
		if !isPrime(i) {
			divisorsCount := countDivisors(i)
			if isPrime(divisorsCount) {
				result = append(result, i)
			}
		}
	}

	return len(result)
}
