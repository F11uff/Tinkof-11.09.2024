package main

import (
	"testing"
)

func TestSimpleNumbers(t *testing.T) {
	tests := []struct {
		l      int64
		r      int64
		result int64
	}{
		{1, 9, 2},
		{10, 20, 1},
		{20, 30, 1},
		{30, 40, 0},
		{40, 50, 1},
		{50, 60, 0},
		{60, 70, 1},
		{70, 80, 0},
		{80, 90, 1},
		{90, 100, 0},
		{1, 1000000, 189},
		{1, 1000000000, 3465},
	}

	for _, tt := range tests {
		result := simpleNumbers(tt.l, tt.r)
		if result != tt.result {
			t.Errorf("ERROR! Expected %d, got %d", tt.result, result)
		}
	}
}
