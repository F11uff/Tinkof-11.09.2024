package main

func main() {
}

func restoreSnowFallData(data []int) (result []int, check bool) {
	dp, result := make([]int, len(data)), make([]int, 0, len(data))
	n := len(data)
	check = true

	if data[0] == -1 {
		data[0] = 1
	}

	for i := 0; i < len(data); i++ {
		if data[i] != -1 {
			dp[i] = data[i]
		} else if i > 0 {
			dp[i] = dp[i-1] + 1
		}
	}

	check = checkError(n, dp)

	if check {
		result = append(result, dp[0])
		for i := 1; i < len(data); i++ {
			result = append(result, dp[i]-dp[i-1])
		}

	}

	return result, check
}

func checkError(n int, dp []int) (check bool) {
	maxRes := dp[0]
	check = true

	if maxRes == 0 {
		check = false
	}

	for i := 1; i < n && check; i++ {
		if maxRes >= dp[i] {
			check = false
			break
		} else {
			maxRes = dp[i]
		}
	}

	return
}
