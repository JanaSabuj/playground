package main

import "testing"

func TestCalc(t *testing.T) {
	if add(15, 15) != 30 {
		t.Error("Add func() is broken!!!")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 4},
		{5, 5, 10},
		{3, 3, 6},
		{6, 6, 12},
	}

	for _, test := range tests {
		if output := add(test.x, test.y); output != test.expected {
			t.Error("Test failed: {} x, {} y , {} expected, {} received", test.x, test.y, test.expected, output)
		}
	}
}
