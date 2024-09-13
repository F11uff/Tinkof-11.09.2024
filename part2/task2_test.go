package main

import (
	"reflect"
	"testing"
)

func TestRestoreSnowFallDataSuccess(t *testing.T) {
	testTrue := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 3, -1, 10, -1}, []int{1, 2, 1, 6, 1}},
		{[]int{3, -1, 6, -1, -1, 20}, []int{3, 1, 2, 1, 1, 12}},
		{[]int{-1, -1, -1, -1, -1, -1}, []int{1, 1, 1, 1, 1, 1}},
		{[]int{5}, []int{5}},
		{[]int{-1}, []int{1}},
		{[]int{200, -1, -1, 300, 301, -1, -1, 500}, []int{200, 1, 1, 98, 1, 1, 1, 197}},
	}

	for _, test := range testTrue {
		resArray, check := restoreSnowFallData(test.input)
		if !reflect.DeepEqual(resArray, test.output) || !check {
			t.Errorf("Input : '%v', output : %v but func output : %v", test.input, test.output, resArray)
		}
	}
}

func TestRestoreSnowFallDataFailed(t *testing.T) {
	testFailed := []struct {
		input []int
	}{
		{[]int{1, 2, 3, 2, 5, 7}},
		{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2}},
		{[]int{100, 1, 2, 3, 4, 5, 6, 7}},
		{[]int{0, 1, 2}},
		{[]int{-1, -2, -3, -4}},
		{[]int{12, 13, 14, -1, -1, -1, -1, 9}},
	}

	for _, test := range testFailed {
		_, check := restoreSnowFallData(test.input)
		if check {
			t.Errorf("Flag check no availabel")
		}
	}
}
