package main

import "fmt"

func main() {
	lastCombination, requirements := "abababababcaba", "ca"
	length := 3
	str := searchPassword(lastCombination, requirements, length)
	fmt.Println(str)
}

func searchPassword(lastCombination, requirements string, length int) string {
	if len(lastCombination) < len(requirements) || len(lastCombination) < length {
		return "-1"
	}

	passwordMaps := make(map[rune]int, length)

	for _, words := range requirements {
		passwordMaps[words] = 0
	}

	fmt.Println("-----", passwordMaps)

	for i := len(lastCombination) - 1; i >= length-1; i-- {
		str := lastCombination[i-length+1 : i+1]
		fmt.Println(str)
		check := true
		for _, words := range str {
			passwordMaps[words]++
		}

		fmt.Println(passwordMaps)

		for _, count := range passwordMaps {
			if count <= 0 {
				check = false
			}
		}

		if check {
			return lastCombination[i-length+1 : i+1]
		}

		passwordMaps = refreshMap(passwordMaps)
	}

	return "-1"
}

func refreshMap(maps map[rune]int) map[rune]int {
	newMap := make(map[rune]int, len(maps))

	for index, _ := range maps {
		newMap[index] = 0
	}

	return newMap
}
