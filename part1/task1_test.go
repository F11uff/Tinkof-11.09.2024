package main

import (
	"reflect"
	"testing"
)

func TestParseRange(t *testing.T) {
	test := []struct {
		input  string
		output []int
	}{
		{"1-6,8-9,11", []int{1, 2, 3, 4, 5, 6, 8, 9, 11}},
		{"1-3,5", []int{1, 2, 3, 5}},
		{"7-10", []int{7, 8, 9, 10}},
		{"4", []int{4}},
		{"10-15,20", []int{10, 11, 12, 13, 14, 15, 20}},
		{"1-1", []int{1}},
		{"5-5", []int{5}},
		{"0-2,4-6", []int{0, 1, 2, 4, 5, 6}},
		{"1,2,3,4,5", []int{1, 2, 3, 4, 5}},
		{"50-55,57,60-62", []int{50, 51, 52, 53, 54, 55, 57, 60, 61, 62}},
		{"100-102,104-106", []int{100, 101, 102, 104, 105, 106}},
		{"1,3,5,7,9", []int{1, 3, 5, 7, 9}},
		{"2,4,6,8,10", []int{2, 4, 6, 8, 10}},
		{"1-2,4-5,7-8", []int{1, 2, 4, 5, 7, 8}},
		{"1,3-5,7-9", []int{1, 3, 4, 5, 7, 8, 9}},
		{"5-7,9-11,13", []int{5, 6, 7, 9, 10, 11, 13}},
		{"100-105", []int{100, 101, 102, 103, 104, 105}},
		{"99", []int{99}},
		{"200-202,204,206-208", []int{200, 201, 202, 204, 206, 207, 208}},
	}

	for _, tt := range test {
		result := parseRange(tt.input)
		if !reflect.DeepEqual(result, tt.output) {
			t.Errorf("Input : '%s', output : %v but func output : %v", tt.input, tt.output, result)
		}
	}
}