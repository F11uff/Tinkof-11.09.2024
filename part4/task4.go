package main

const SQRT = 1e7

func main() {

}

func simpleNumbers(l, r int64) int64 {
	tableEratosfens := sieveOfEratosthenes(int(SQRT))

	return countNumbers(l, r, tableEratosfens)
}

func sieveOfEratosthenes(countOfMany int) []int {
	flagsForTableEratosfens := make([]bool, countOfMany+1)

	for i := range flagsForTableEratosfens {
		flagsForTableEratosfens[i] = true
	}

	flagsForTableEratosfens[0], flagsForTableEratosfens[1] = false, false

	tableEratosfens := make([]int, 0)

	for i := 2; i <= countOfMany; i++ {
		if flagsForTableEratosfens[i] {
			tableEratosfens = append(tableEratosfens, i)
			for j := 2 * i; j <= countOfMany; j += i {
				flagsForTableEratosfens[j] = false
			}
		}
	}

	return tableEratosfens
}

func countNumbers(l, r int64, tableEratosfens []int) int64 {
	ans := int64(0)

	for _, value := range tableEratosfens {
		x := int64(value)

		if x > r {
			break
		}

		degIn := countDivisors(r, x)
		degOut := countDivisors(l-1, x)

		for deg := degOut + 1; deg <= degIn; deg++ {
			if deg > 1 && isPrime(deg+1, tableEratosfens) {
				ans++
			}
		}
	}

	return ans
}

func countDivisors(n int64, value int64) int {
	deg := 0

	for n/value > 0 {
		deg++
		n /= value
	}

	return deg
}

func isPrime(n int, tableEratosfens []int) bool {
	if n < 2 {
		return false
	}

	for _, p := range tableEratosfens {
		if p*p > n {
			break
		}

		if n%p == 0 {
			return false
		}
	}

	return true
}
