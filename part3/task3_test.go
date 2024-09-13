package main

import (
	"reflect"
	"testing"
)

func TestSearchPassword(t *testing.T) {
	test := []struct {
		lastCombination string
		requirements    string
		length          int
		output          string
	}{
		{"abacaba", "abc", 4, "caba"},
		{"abacaba", "abc", 3, "cab"},
		{"abcdddcbadacbdddbcbbasb", "abcd", 5, "bcbba"},
		{"basdfgrtdheyaaaaaag", "a", 7, "-1"},
		{"basdfgrtdheyaaaaaaag", "a", 7, "aaaaaaa"},
		{"qwewecweqeceewcqww", "qwe", 5, "qwewe"},
		{"12453456745654", "76", 1, "-1"},
		{"124534456745654", "45", 3, "445"},
		{"abababababcaba", "ca", 3, "-1"},
		{"qwwqqwwwqwwwqwwqqqweww", "qw", 5, "wqqqw"},
		{"qww", "qw", 5, "-1"},
	}

	for _, tt := range test {
		outputStr := searchPassword(tt.lastCombination, tt.requirements, tt.length)
		if !reflect.DeepEqual(outputStr, tt.output) {
			t.Errorf("Input : '%v', output : %v but func output : %v", tt.lastCombination, tt.output, outputStr)
		}
	}
}
