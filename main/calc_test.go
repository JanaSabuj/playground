package main

import "testing"

func TestCalc(t *testing.T) {
	if add(15, 15) != 30 {
		t.Error("Add func() is broken!!!")
	}
}
