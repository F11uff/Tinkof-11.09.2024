package main

import (
	"strings"
)

func main() {

}

func searchPassword(lastCombination, requirements string, length int) string {
	if len(lastCombination) < len(requirements) || length < len(requirements) {
		return "-1"
	}

	passwordMaps := make(map[rune]int, length)

	for _, words := range requirements {
		passwordMaps[words] = 0
	}

	for i := len(lastCombination) - 1; i >= length-1; i-- {
		str := lastCombination[i-length+1 : i+1]
		check := true

		for _, words := range str {
			passwordMaps[words]++
		}

		for index, count := range passwordMaps {
			if count <= 0 || !strings.Contains(requirements, string(index)) {
				check = false
			}
		}

		if check {
			return str
		}

		passwordMaps = refreshMap(passwordMaps)
	}

	return "-1"
}

func refreshMap(maps map[rune]int) map[rune]int {
	newMap := make(map[rune]int, len(maps))

	for index := range maps {
		delete(maps, index)
	}

	return newMap
}
