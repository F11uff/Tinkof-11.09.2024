package main

import (
	"reflect"
	"testing"
)

func TestSearchPassword(t *testing.T) {
	testSuccess := []struct {
		lastCombination string
		requirements    string
		length          int
		output          string
	}{
		//{"abacaba", "abc", 4, "caba"},
		//{"abacaba", "abc", 3, "cab"},
		{"abababababcaba", "ca", 3, "-1"},
	}

	for _, tt := range testSuccess {
		outputStr := searchPassword(tt.lastCombination, tt.requirements, tt.length)
		if !reflect.DeepEqual(outputStr, tt.output) {
			t.Errorf("Input : '%v', output : %v but func output : %v", tt.lastCombination, tt.output, outputStr)
		}
	}
}
