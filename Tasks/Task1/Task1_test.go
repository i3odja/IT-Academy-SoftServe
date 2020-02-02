package main

import (
	"testing"
)

/*
func TestMultiplicity3not5(t *testing.T) {
	a := []int{3, 3, 15}
	rez := Multiplicity3not5(a)
	if rez != 2 {
		t.Errorf("Count is incorrect: %d, want: %d.", rez, 2)
	}
}*/
type TestCase struct {
	input    []int
	expCount int
}

func TestMultiplicity3not5(t *testing.T) {
	b := TestCase{
		[]int{3, 4, 4, 3, 9, 15},
		3,
	}
	//a := []int{3, 3, 15}
	rez := Multiplicity3not5(b.input)
	if rez != b.expCount {
		t.Errorf("Count is incorrect: %d, want: %d.", rez, b.expCount)
	}
}
