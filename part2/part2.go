package main

//5
//1 3 -1 10 -1
//
//YES
//1 2 3 4 5

func main() {
}

func restoreSnowFallData(data []int) (result []int, check bool) {
	res, result := make([]int, len(data)), make([]int, 0, len(data))
	n := len(data)
	check = true

	if data[0] == -1 {
		data[0] = 1
	}

	for i := 0; i < len(data); i++ {
		if data[i] != -1 {
			res[i] = data[i]
		} else if i > 0 {
			res[i] = res[i-1] + 1
		}
	}

	check = checkError(n, res)

	if check {
		result = append(result, res[0])
		for i := 1; i < len(data); i++ {
			result = append(result, res[i]-res[i-1])
		}

	}

	return result, check
}

func checkError(n int, res []int) (check bool) {
	maxRes := res[0]
	check = true

	if maxRes == 0 {
		check = false
	}

	for i := 1; i < n && check; i++ {
		if maxRes >= res[i] {
			check = false
			break
		} else {
			maxRes = res[i]
		}
	}

	return
}
