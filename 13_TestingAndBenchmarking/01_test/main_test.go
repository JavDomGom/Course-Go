package main

import "testing"

func TestMiSuma(t *testing.T) {
	v := miSuma(2, 8)
	if v != 10 {
		t.Error("Expected", 10, ", got", miSuma(2, 8))
	}
}
