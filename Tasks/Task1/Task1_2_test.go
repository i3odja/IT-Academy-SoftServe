package main

import (
	"testing"
)

type TestCase2 struct {
	input []int
	expCount int
	msg string
}

func TestMultiplicity3not5_v2(t *testing.T) {
	b := []TestCase2{
		TestCase2{
			[]int{3, 3, 3, 9, 9, 27, 33, 36, 12},
			9,
			"All items are correct",
		},
		TestCase2{
			[]int{30, 60, 15, 45, 90, 120, 180},
			0,
			"No one item is correct",
		},
	}
	for _, v := range b {
		rez := Multiplicity3not5(v.input)
		if rez != v.expCount {
			t.Errorf("Count is incorrect: %d, want: %d. %s", rez, v.expCount, v.msg)
		}
	}
}
