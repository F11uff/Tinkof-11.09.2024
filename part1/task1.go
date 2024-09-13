package main

import (
	"strconv"
	"strings"
)

func main() {
}

func parseRange(s string) []int {
	split := strings.Split(s, ",")
	res := []int{}

	for _, val := range split {
		if strings.Contains(val, "-") {
			split_ := strings.Split(val, "-")
			start, _ := strconv.Atoi(split_[0])
			end, _ := strconv.Atoi(split_[1])
			for i := start; i <= end; i++ {
				res = append(res, i)
			}
		} else {
			count, _ := strconv.Atoi(val)
			res = append(res, count)
		}
	}

	return res
}
