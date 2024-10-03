package main

import "testing"

func TestSubstract(t *testing.T) {
	total := substract(5, 5)
	if total != 0 {
		t.Errorf("Output should be %d but total is %v.", 0, total)
	}
}
